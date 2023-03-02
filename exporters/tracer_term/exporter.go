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

// Package tracer_term
// 链路[同步]打印到 终端/控制台.
package tracer_term

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/formatters/tracer_term"
	"github.com/fuyibing/log/v5/traces"
	"github.com/fuyibing/util/v8/process"
	"os"
)

type Exporter struct {
	formatter  traces.TracerFormatter
	name       string
	processing int32
	processor  process.Processor
}

func New() traces.TracerExporter { return (&Exporter{}).init() }

// Processor 获取类进程.
func (o *Exporter) Processor() process.Processor { return o.processor }

// Send 发送日志.
func (o *Exporter) Send(span traces.Span) error {
	_, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", o.format(span)))
	return err
}

// SetFormatter 设置格式化.
func (o *Exporter) SetFormatter(formatter traces.TracerFormatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) format(v traces.Span) string {
	if o.formatter != nil {
		return o.formatter.String(v)
	}

	return fmt.Sprintf("# [%s] %s",
		v.GetSpanId().String(),
		v.GetName(),
	)
}

func (o *Exporter) init() *Exporter {
	o.formatter = tracer_term.NewFormatter()

	o.name = "exporter.tracer.term"
	o.processor = process.New(o.name).
		Callback(o.onCall).
		Panic(o.onPanic)
	return o
}

func (o *Exporter) onCall(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *Exporter) onPanic(_ context.Context, v interface{}) {
	traces.InternalError("<%s> %v", o.name, v)
}
