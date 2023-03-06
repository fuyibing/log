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
// date: 2023-03-05

package tracer_term

import (
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/tracers"
	"strings"
)

var colors = map[common.Level][]int{
	common.Debug: {37, 0},  // Text: gray,   Background: x
	common.Info:  {34, 0},  // Text: blue,   Background: x
	common.Warn:  {33, 0},  // Text: yellow, Background: x
	common.Error: {31, 0},  // Text: red,    Background: x
	common.Fatal: {31, 43}, // Text: red,    Background: yellow
}

type formatter struct{}

func (o *formatter) Byte(_ ...tracers.Span) (body []byte, err error) { return }

func (o *formatter) String(vs ...tracers.Span) (text string, err error) {
	var list = make([]string, 0)
	for _, v := range vs {
		list = append(list, o.format(v))
	}

	text = strings.Join(list, "\n")
	return
}

func (o *formatter) format(v tracers.Span) (text string) {
	var (
		sid = v.SpanId().String()
		pid = v.ParentSpanId().String()
		tid = v.Trace().TraceId().String()
	)

	// 跨度信息.
	text = fmt.Sprintf("[span=%s][parent=%s][trace=%s] duration=%v us, %s",
		sid, pid, tid,
		v.Duration().Microseconds(), v.Name(),
	)

	// 跨度属性.
	if kv := v.Kv(); len(kv) > 0 {
		text += fmt.Sprintf("\n[span=%s][$] %s", sid, kv.String())
	}

	// 跨度日志.
	for _, log := range v.Logs() {
		// 记录时间.
		text += fmt.Sprintf("\n[span=%s][%-15s]", sid, log.Time().Format("15:04:05.999999"))

		// 日志级别.
		if c, ok := colors[log.Level()]; ok {
			// 着色.
			text += fmt.Sprintf("[%s]",
				fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
					0x1B, 0, c[1], c[0], fmt.Sprintf("%5s", log.Level()), 0x1B,
				),
			)
		} else {
			// 无色.
			text += fmt.Sprintf("[%5s]", log.Level())
		}

		// 日志键值对.
		if kv := log.Kv(); len(kv) > 0 {
			text += fmt.Sprintf(" %s", kv.String())
		}

		// 日志正文.
		text += fmt.Sprintf(" %s", log.Text())

		// 异常堆栈.
		if log.Stack() {
			for _, item := range log.Stacks() {
				if item.Internal {
					continue
				}
				text += fmt.Sprintf("\n[span=%s][#] <%s:%d> IN <%s>", sid,
					item.File,
					item.Line,
					item.Call,
				)
			}
		}
	}
	return
}

func (o *formatter) init() *formatter {
	return o
}
