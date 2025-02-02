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
	"fmt"
	"io"
)

func handleHistory(rc io.ReadCloser) error {
	decoder := json.NewDecoder(rc)
	for decoder.More() {
		var history history
		if err := decoder.Decode(&history); err != nil {
			return fmt.Errorf("failed to parse history body data %s", err.Error())
		}

		errors := history.validate()
		if len(errors) != 0 {
			return fmt.Errorf("failed to validate data errors: %v", errors)
		}
	}

	return nil
}

func (h history) validate() map[string]string {
	errors := h.generic.validate()

	err := h.Type.validate()
	if err != nil {
		errors["type"] = err.Error()
	}

	err = h.Itemid.validate()
	if err != nil {
		errors["itemid"] = err.Error()
	}

	err = h.Host.validate()
	if err != nil {
		errors["host"] = err.Error()
	}

	err = h.Tags.validate()
	if err != nil {
		errors["item_tags"] = err.Error()
	}

	for i, t := range h.Tags.Value {
		err = t.Tag.validate()
		if err != nil {
			errors[fmt.Sprintf("item_tags[%d].host", i)] = err.Error()
		}

		err = t.Value.validate()
		if err != nil {
			errors[fmt.Sprintf("item_tags[%d].name", i)] = err.Error()
		}
	}

	return errors
}
