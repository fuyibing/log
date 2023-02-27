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
// date: 2023-02-27

package logger_term

import (
	"fmt"
	"github.com/fuyibing/log/v5/cores"
)

type (
	// Formatter
	// 格式化接口.
	Formatter interface {
		// Format
		// 格式化.
		Format(line cores.Line) (text string)

		// SetColor
		// 设置着色.
		SetColor(color bool) Formatter

		// SetTimeFormat
		// 设置时间格式.
		SetTimeFormat(format string) Formatter
	}

	formatter struct {
		// 是否着色.
		// 在 Terminal/Console 输出日志时是否着色.
		Color bool

		// 时间格式.
		TimeFormat string
	}
)

// Format
// 格式化.
func (o *formatter) Format(line cores.Line) (text string) {
	// 日志属性.
	attr := ""
	if buf, err := line.GetAttr().Marshal(); err == nil {
		if attr = string(buf); attr != "" {
			attr += " "
		}
	}

	// 日志正文.
	text = fmt.Sprintf("[%-15s][%5v] %s%s",
		line.GetTime().Format(o.TimeFormat),
		line.GetLevel(), attr, line.GetText(),
	)

	// 日志着色.
	if o.Color {
		if c, ok := colors[line.GetLevel()]; ok {
			text = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
				0x1B, 0, c[1], c[0], text, 0x1B,
			)
		}
	}

	return
}

// SetColor
// 设置着色.
func (o *formatter) SetColor(color bool) Formatter {
	o.Color = color
	return o
}

// SetTimeFormat
// 设置时间格式.
func (o *formatter) SetTimeFormat(format string) Formatter {
	o.TimeFormat = format
	return o
}

func (o *formatter) init() *formatter {
	o.Color = defaultColor
	o.TimeFormat = defaultTimeFormat
	return o
}
