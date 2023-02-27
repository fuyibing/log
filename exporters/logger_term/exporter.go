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
// date: 2023-02-25

package logger_term

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/util/v8/process"
	"os"
	"sync/atomic"
	"time"
)

var (
	colors = map[base.Level][]int{
		base.Debug: {37, 0},  // Text: gray, Background: white
		base.Info:  {34, 0},  // Text: blue, Background: white
		base.Warn:  {33, 0},  // Text: yellow, Background: white
		base.Error: {31, 0},  // Text: red, Background: white
		base.Fatal: {33, 41}, // Text: yellow, Background: red
	}
)

const (
	defaultColor      = true
	defaultTimeFormat = "15:04:05.999999"
)

type (
	// Exporter
	// Log 输出到 Terminal/Console.
	Exporter struct {
		formatter  Formatter
		processing int32
		processor  process.Processor
	}
)

func New() *Exporter { return (&Exporter{}).init() }

// Processor
// 执行器.
func (o *Exporter) Processor() process.Processor { return o.processor }

// Push
// 输出日志.
func (o *Exporter) Push(logs ...cores.Line) {
	atomic.AddInt32(&o.processing, 1)
	defer atomic.AddInt32(&o.processing, -1)

	// 遍历日志.
	for _, log := range logs {
		_, _ = fmt.Fprintf(os.Stdout,
			fmt.Sprintf("%s\n",
				o.formatter.Format(log),
			),
		)
	}
}

// SetFormatter
// 设置格式化.
func (o *Exporter) SetFormatter(formatter Formatter) *Exporter {
	o.formatter = formatter
	return o
}

func (o *Exporter) init() *Exporter {
	o.formatter = (&formatter{}).init()
	o.processor = process.New("term-logger-exporter").
		After(o.onAfter).
		Callback(o.onCall)
	return o
}

func (o *Exporter) onAfter(ctx context.Context) (ignored bool) {
	if atomic.LoadInt32(&o.processing) == 0 {
		return
	}
	time.Sleep(time.Millisecond)
	return o.onAfter(ctx)
}

func (o *Exporter) onCall(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *Exporter) onPanic(_ context.Context, _ interface{}) {
	cores.Registry.Debugger("%s panic: %v", o.processor.Name())
}
