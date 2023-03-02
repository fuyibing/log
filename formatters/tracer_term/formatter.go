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

package tracer_term

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"strings"
)

type (
	// Formatter
	// 终端格式化.
	Formatter struct{}
)

func NewFormatter() base.TracerFormatter { return &Formatter{} }

// Byte
// 转成Byte切片.
func (o *Formatter) Byte(_ base.Span) []byte { return nil }

// String
// 转成字符串.
func (o *Formatter) String(v base.Span) string {
	// 基础信息.
	//
	// - # [SID=string][PID=string][TID=string] duration=37 us
	var list = []string{
		fmt.Sprintf("$ [SID=%s][PID=%s][TID=%s] %v us, %v",
			v.GetSpanId().String(),
			v.GetParentSpanId().String(),
			v.GetTrace().GetTraceId(),
			v.GetDuration().Microseconds(),
			v.GetName(),
		),
	}

	// 属性信息.
	if v.GetAttribute().Count() > 0 {
		list = append(list, fmt.Sprintf("  ! %s", v.GetAttribute().String()))
	}

	// 日志信息.
	for _, log := range v.GetLogs() {
		list = append(list, fmt.Sprintf("  # [%-15s][%5s] %s",
			log.GetTime().Format("15:04:05.999999"),
			log.GetLevel(),
			log.GetText(),
		))
	}

	return strings.Join(list, "\n")
}
