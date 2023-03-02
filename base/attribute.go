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
// date: 2023-03-01

package base

import (
	"encoding/json"
)

type (
	// Attribute 跟踪属性.
	Attribute map[string]interface{}
)

// Add 添加 Key/Value 键值对.
func (o Attribute) Add(key string, value interface{}) Attribute {
	o[key] = value
	return o
}

// Copy 复制 Key/Value 键值对.
func (o Attribute) Copy(src Attribute) Attribute {
	for key, value := range src {
		o[key] = value
	}
	return o
}

// Count 元素数量.
func (o Attribute) Count() int {
	return len(o)
}

// Marshal 转为JSON字符码.
func (o Attribute) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

// String 返回JSON字符串.
func (o Attribute) String() (str string) {
	if buf, err := o.Marshal(); err == nil {
		str = string(buf)
	}
	return
}
