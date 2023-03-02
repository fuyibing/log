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

package logger_term

import (
	"fmt"
	"github.com/fuyibing/log/v5/traces"
)

var (
	// 终端着色.
	termColor = map[traces.Level][]int{
		traces.Debug: {37, 0},  // Text: gray, Background: white
		traces.Info:  {34, 0},  // Text: blue, Background: white
		traces.Warn:  {33, 0},  // Text: yellow, Background: white
		traces.Error: {31, 0},  // Text: red, Background: white
		traces.Fatal: {33, 41}, // Text: yellow, Background: red
	}
)

type (
	// Formatter
	// 终端格式化.
	Formatter struct{}
)

func NewFormatter() traces.LoggerFormatter { return &Formatter{} }

// Byte
// 转成Byte切片.
func (o *Formatter) Byte(_ traces.Log) []byte { return nil }

// String
// 转成字符串.
func (o *Formatter) String(v traces.Log) (text string) {
	// 基础信息.
	// - 时间
	// - 级别
	text = fmt.Sprintf("[%-15s][%5s]",
		v.GetTime().Format("15:04:05.999999"),
		v.GetLevel(),
	)

	// 日志属性.
	if v.GetAttribute().Count() > 0 {
		text += fmt.Sprintf(" %s", v.GetAttribute().String())
	}

	// 日志正文.
	text += fmt.Sprintf(" %s", v.GetText())

	// 日志着色.
	if c, ok := termColor[v.GetLevel()]; ok {
		text = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
			0x1B,
			0,
			c[1],
			c[0],
			text,
			0x1B,
		)
	}

	return
}
