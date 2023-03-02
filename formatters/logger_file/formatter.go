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

package logger_file

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
)

type (
	// Formatter
	// 文本格式化.
	Formatter struct{}
)

func NewFormatter() base.LoggerFormatter { return &Formatter{} }

// Byte
// 转成Byte切片.
func (o *Formatter) Byte(_ base.Log) []byte { return nil }

// String
// 转成字符串.
func (o *Formatter) String(v base.Log) (text string) {
	// 基础信息.
	// - 时间
	// - 级别
	text = fmt.Sprintf("[%-26s][%5s]",
		v.GetTime().Format("2006-01-02 15:04:05.999999"),
		v.GetLevel(),
	)

	// 日志属性.
	if v.GetAttribute().Count() > 0 {
		text += fmt.Sprintf(" %s", v.GetAttribute().String())
	}

	// 日志正文.
	text += fmt.Sprintf(" %s", v.GetText())
	return
}
