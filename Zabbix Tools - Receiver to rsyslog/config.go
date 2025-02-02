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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var defaultConfigPath = "./config.json"

type Config struct {
	EnableTls bool   `json:"enable_tls"`
	Port      string `json:"port"`
	DataPath  string `json:"data_path"`
	LogPath   string `json:"log_path"`
	CertFile  string `json:"cert_file"`
	KeyFile   string `json:"key_file"`
}

func (c *Config) setDefaults() error {
	if c == nil {
		return errors.New("config not set")
	}

	if c.Port == "" {
		c.Port = "80"
	}

	fp, err := filepath.Abs(os.Args[0])
	if err != nil {
		return fmt.Errorf("failed to find executable location, %s", err.Error())
	}

	if c.DataPath == "" {
		c.DataPath = filepath.Join(filepath.Dir(fp), "data")
	}

	if c.LogPath == "" {
		c.LogPath = fmt.Sprintf("%s.log", fp)
	}

	return nil
}

func loadConfiguration(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s", err.Error())
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %s", err.Error())
	}

	config.setDefaults()

	return &config, nil
}
