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
	// 链路属性组件, 用于 Trace, Span, Line 组件.
	Attr interface {
		// Add
		// 向本组件添加KV键值对.
		Add(key string, value interface{}) Attr

		// Copy
		// 复制源组件的KV键值对, 并加到本组件.
		Copy(s Attr) Attr

		// GetMap
		// 返回本组件KV键值对的Map结构.
		GetMap() map[string]interface{}

		// Marshal
		// 转成JSON格式.
		Marshal() ([]byte, error)

		// With
		// 绑定任意类型数据到KV键值对.
		With(v interface{}) Attr
	}

	attr struct {
		sync.RWMutex
		data map[string]interface{}
	}
)

// NewAttr
// 创建链路属性组件.
func NewAttr() Attr {
	return &attr{
		data: make(map[string]interface{}),
	}
}

// Add
// 向本组件添加KV键值对.
func (o *attr) Add(key string, value interface{}) Attr {
	o.Lock()
	defer o.Unlock()

	o.data[key] = value
	return o
}

// Copy
// 复制源组件的KV键值对, 并加到本组件.
func (o *attr) Copy(s Attr) Attr {
	o.Lock()
	defer o.Unlock()

	for key, value := range s.GetMap() {
		o.data[key] = value
	}
	return o
}

// GetMap
// 返回本组件KV键值对的Map结构.
func (o *attr) GetMap() map[string]interface{} {
	o.RLock()
	defer o.RUnlock()

	return o.data
}

// Marshal
// 转成JSON格式.
func (o *attr) Marshal() (buf []byte, err error) {
	o.RLock()
	defer o.RUnlock()

	if len(o.data) > 0 {
		buf, err = json.Marshal(o.data)
	}
	return
}

// With
// 绑定任意类型数据到KV键值对.
func (o *attr) With(data interface{}) Attr {
	if data == nil {
		return o
	}

	// 同类型复杂.
	if m, ok := data.(Attr); ok {
		return o.Copy(m)
	}

	// Map映射.
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
