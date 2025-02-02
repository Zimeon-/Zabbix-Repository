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
)

type generic struct {
	Ns     optional[int]      `json:"ns"`
	Clock  optional[int]      `json:"clock"`
	Value  optional[float64]  `json:"value"`
	Name   optional[string]   `json:"name"`
	Groups optional[[]string] `json:"groups"`
}

type event struct {
	generic
	EventId  optional[int]    `json:"eventid"`
	Severity optional[int]    `json:"severity"`
	Hosts    optional[[]host] `json:"hosts"`
	Tags     optional[[]tag]  `json:"tags"`
}

type history struct {
	generic
	Type   optional[int]   `json:"type"`
	Itemid optional[int]   `json:"itemid"`
	Host   optional[host]  `json:"host"`
	Tags   optional[[]tag] `json:"item_tags"`
}

type host struct {
	Host optional[string] `json:"host"`
	Name optional[string] `json:"name"`
}

type tag struct {
	Tag   optional[string] `json:"tag"`
	Value optional[string] `json:"value"`
}

type optional[T any] struct {
	Defined bool
	Value   T
}

func (o *optional[T]) validate() error {
	if !o.Defined {
		return errors.New("not set")
	}

	return nil
}

func (o *optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.Value)
}
