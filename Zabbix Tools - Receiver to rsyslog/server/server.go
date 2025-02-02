/*
** Copyright (C) 2001-2025 Zabbix SIA
**
** Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
** documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
** rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
** permit persons to whom the Software is furnished to do so, subject to the following conditions:
**
** The above copyright notice and this permission notice shall be included in all copies or substantial portions
** of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
** WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
** COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
** TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
** SOFTWARE.
**/

package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	contentType        = "Content-Type"
	applicationXndJson = "application/x-ndjson"
	applicationJson    = "application/json"
	historyFilename    = "history.ndjson"
	eventsFilename     = "events.ndjson"
)

var savePath string
var fileMux sync.Mutex

type payloadHandler struct {
	filename    string
	handlerFunc func(io.ReadCloser) error
}

var postHandlers = map[string]payloadHandler{
	"/v1/events":  {eventsFilename, handleEvent},
	"/v1/history": {historyFilename, handleHistory},
}

func Run(port, cert, key, dataPath string, tls bool) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	savePath = dataPath

	path := fmt.Sprintf(":%s", port)

	if tls {
		err := validateTLS(cert, key)
		if err != nil {
			return err
		}

		return http.ListenAndServeTLS(path, cert, key, mux)
	}

	return http.ListenAndServe(path, mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	h, ok := postHandlers[r.URL.Path]
	if !ok {
		jsonResponse(w, nil, "404 page not found", http.StatusNotFound)
		return
	}

	postHandler(w, r, h)
}

func postHandler(w http.ResponseWriter, r *http.Request, h payloadHandler) {
	if r.Method != http.MethodPost {
		jsonResponse(
			w,
			nil,
			fmt.Sprintf("method %s is not allowed", r.Method),
			http.StatusMethodNotAllowed,
		)

		return
	}

	if r.Header.Get(contentType) != applicationXndJson {
		jsonResponse(
			w,
			nil,
			fmt.Sprintf("%s header must contain %s", contentType, applicationXndJson),
			http.StatusUnsupportedMediaType,
		)

		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsonResponse(
			w,
			body,
			fmt.Sprintf("failed to read body data %s", err.Error()),
			http.StatusBadRequest,
		)

		return
	}

	err = h.handlerFunc(r.Body)
	if err != nil {
		jsonResponse(w, body, err.Error(), http.StatusBadRequest)

		return
	}

	err = saveData(h.filename, body)
	if err != nil {
		jsonResponse(w, body, err.Error(), http.StatusInternalServerError)

		return
	}

	jsonResponse(w, body, "", http.StatusOK)
}

func saveData(filename string, body []byte) error {
	fileMux.Lock()
	defer fileMux.Unlock()
	file, err := os.OpenFile(filepath.Join(savePath, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
			return fmt.Errorf("failed to open save file %s", err.Error())
	}

	defer file.Close()

	//_, err = file.WriteString(string(body))
	//if err != nil {
	//      return fmt.Errorf("failed to write to save file %s", err.Error())
	//}

	// Send log data to rsyslog server
	if err := sendToSyslog(string(body)); err != nil {
			log.Printf("Warning: Failed to send log to syslog: %v", err)
	}

	return nil
}

// sendToSyslog sends logs to a remote rsyslog server
func sendToSyslog(message string) error {
	sysLogger, err := syslog.Dial("udp", "127.0.0.1:514", syslog.LOG_INFO|syslog.LOG_USER, "ZbxStream")
	if err != nil {
			return fmt.Errorf("failed to connect to syslog: %w", err)
	}
	defer sysLogger.Close()

	return sysLogger.Info(message)
}

func (h generic) validate() map[string]string {
	errors := make(map[string]string)

	err := h.Clock.validate()
	if err != nil {
		errors["clock"] = err.Error()
	}

	err = h.Ns.validate()
	if err != nil {
		errors["ns"] = err.Error()
	}

	err = h.Value.validate()
	if err != nil {
		errors["value"] = err.Error()
	}

	err = h.Name.validate()
	if err != nil {
		errors["name"] = err.Error()
	}

	err = h.Groups.validate()
	if err != nil {
		errors["groups"] = err.Error()
	}

	return errors
}

func jsonResponse(w http.ResponseWriter, body []byte, msg string, code int) {
	w.Header().Set("Content-Type", applicationJson)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if code == http.StatusOK {
		log.Printf("request successful with data: %s", body)
		j, _ := json.Marshal(map[string]string{"response": "success"})
		fmt.Fprintln(w, string(j))
		return
	}

	if body == nil {
		log.Printf("request failed without data: %s", msg)
	} else {
		log.Printf("request failed with data: %s, error: %s", body, msg)
	}

	j, _ := json.Marshal(map[string]string{"response": "fail", "info": msg})
	fmt.Fprintln(w, string(j))
}

func validateTLS(certPath, keyPath string) error {
	if certPath == "" || keyPath == "" {
		return errors.New("both tls certificate and key file paths must be set")
	}

	return nil
}
