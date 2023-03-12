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
// date: 2023-03-04

package loggers

import (
	"encoding/json"
)

type (
	// Kv
	// component for logger, stored as key/value.
	Kv map[string]interface{}
)

func (o Kv) Add(key string, value interface{}) Kv {
	o[key] = value
	return o
}

func (o Kv) Copy(s Kv) Kv {
	for k, v := range s {
		o[k] = v
	}
	return o
}

func (o Kv) String() (str string) {
	if o != nil {
		if buf, err := json.Marshal(o); err == nil {
			str = string(buf)
		}
	}
	return
}
