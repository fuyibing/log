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

// Package logger_term
// 日志[同步]打印到 终端/控制台.
package logger_term

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/formatters/logger_term"
	"github.com/fuyibing/util/v8/process"
	"os"
)

type Exporter struct {
	formatter base.LoggerFormatter
	name      string
	processor process.Processor
}

func New() base.LoggerExporter { return (&Exporter{}).init() }

// Processor 获取类进程.
func (o *Exporter) Processor() process.Processor { return o.processor }

// Send 发送日志.
func (o *Exporter) Send(log base.Log) error {
	_, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", o.formatter.String(log)))
	return err
}

// SetFormatter 设置格式化.
func (o *Exporter) SetFormatter(formatter base.LoggerFormatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) format(log base.Log) string {
	if o.formatter != nil {
		return o.formatter.String(log)
	}

	return fmt.Sprintf("[%-15s][%5s] %s",
		log.GetTime().Format("15:04:05.999999"),
		log.GetLevel(), log.GetText(),
	)
}

func (o *Exporter) init() *Exporter {
	o.formatter = logger_term.NewFormatter()

	o.name = "exporter.logger.term"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall).
		Panic(o.onPanic)
	return o
}

func (o *Exporter) onAfter(_ context.Context) (ignored bool) {
	return
}

func (o *Exporter) onBefore(_ context.Context) (ignored bool) {
	return
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
	base.InternalError("<%s> %v", o.name, v)
}
