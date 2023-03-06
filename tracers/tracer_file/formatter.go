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

package tracer_file

import (
	"fmt"
	"github.com/fuyibing/log/v5/tracers"
	"strings"
)

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
		// 时间与级别.
		text += fmt.Sprintf("\n[span=%s][%-26s][%5s]", sid, log.Time().Format("2006-01-02 15:04:05.999999"), log.Level())

		// 日志键值对.
		if kv := log.Kv(); len(kv) > 0 {
			text += fmt.Sprintf(" %s", kv.String())
		}

		// 日志下文.
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

func (o *formatter) init() *formatter { return o }
