// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: wsfuyibing <websearch@163.com>
// date: 2023-02-26

package cores

import (
	"encoding/json"
	"sync"
)

type (
	// Attr
	// is the component for tracer, used on Trace, Span, Line components.
	Attr interface {
		// Add
		// key/value pair into an Attr component.
		Add(key string, value interface{}) Attr

		// Copy
		// key/value pair from a source to an Attr component.
		Copy(s Attr) Attr

		// GetMap
		// returns a map struct of the Attr component.
		GetMap() map[string]interface{}

		// Marshal
		// encode as json byte.
		Marshal() ([]byte, error)

		With(v interface{}) Attr
	}

	attr struct {
		sync.RWMutex
		data map[string]interface{}
	}
)

// NewAttr
// returns an Attr component.
func NewAttr() Attr {
	return &attr{
		data: make(map[string]interface{}),
	}
}

// Add
// key/value pair into an Attr component.
func (o *attr) Add(key string, value interface{}) Attr {
	o.Lock()
	defer o.Unlock()
	o.data[key] = value
	return o
}

// Copy
// key/value pair from a source to an Attr component.
func (o *attr) Copy(s Attr) Attr {
	o.Lock()
	defer o.Unlock()
	for key, value := range s.GetMap() {
		o.data[key] = value
	}
	return o
}

// GetMap
// returns a map struct of the Attr component.
func (o *attr) GetMap() map[string]interface{} {
	o.RLock()
	defer o.RUnlock()
	return o.data
}

// Marshal
// encode as json byte.
func (o *attr) Marshal() (buf []byte, err error) {
	o.RLock()
	defer o.RUnlock()
	if len(o.data) > 0 {
		buf, err = json.Marshal(o.data)
	}
	return
}

func (o *attr) With(data interface{}) Attr {
	if data == nil {
		return o
	}

	// Attr mapping.
	if m, ok := data.(Attr); ok {
		o.Lock()
		defer o.Unlock()
		for k, v := range m.GetMap() {
			o.data[k] = v
		}
		return o
	}

	// Map type.
	if m, ok := data.(map[string]interface{}); ok {
		o.Lock()
		defer o.Unlock()
		for k, v := range m {
			o.data[k] = v
		}
		return o
	}

	return o
}
