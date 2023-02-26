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

package tracer_term

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

type (
	// Exporter
	// an exporter component for tracer, print on terminal/console.
	Exporter struct {
		processing int32
		processor  process.Processor
	}
)

// NewExporter
// returns an exporter component.
func NewExporter() *Exporter {
	return (&Exporter{}).init()
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: interface methods
// /////////////////////////////////////////////////////////////////////////////

// Processor
// returns a processor component.
func (o *Exporter) Processor() process.Processor {
	return o.processor
}

// Push
// log to exporter to do write / publish operation.
func (o *Exporter) Push(spans ...cores.Span) {
	atomic.AddInt32(&o.processing, 1)
	defer atomic.AddInt32(&o.processing, -1)

	for _, span := range spans {
		o.printSpan(span)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: internal methods
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) printSpan(span cores.Span) {
	// Append Span info.
	list := []string{o.withColor(base.Info, o.renderHead(span))}

	// Append source.
	for k, v := range cores.Registry.Resource().GetMap() {
		list = append(list,
			fmt.Sprintf("       - {%11s} : %v", k, v),
		)
	}

	// Append logs.
	for _, log := range span.GetLogs().GetLines() {
		list = append(list,
			"     + "+o.withColor(log.GetLevel(), o.renderLog(log)),
		)

		for k, v := range log.GetAttr().GetMap() {
			list = append(list,
				"       - "+o.withColor(base.Debug,
					fmt.Sprintf("{%s} : %v", k, v),
				),
			)
		}
	}

	// Print lines.
	for _, str := range list {
		_, _ = fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", str))
	}
}

func (o *Exporter) renderHead(span cores.Span) string {
	return fmt.Sprintf("SPAN + [%s][%s][%s] %s",
		span.GetSpanId().String(),
		span.GetParentSpanId().String(),
		span.GetTraceId().String(),
		span.GetName(),
	)
}

func (o *Exporter) renderLog(log cores.Line) string {
	return fmt.Sprintf("[%-16s][%5v] %s",
		log.GetTime().Format("15:04:05.9999999"),
		log.GetLevel(),
		log.GetText(),
	)
}

func (o *Exporter) withColor(level base.Level, content string) string {
	if c, ok := colors[level]; ok {
		return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
			0x1B, 0, c[1], c[0], content, 0x1B,
		)
	}
	return content
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: events
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) onAfter(ctx context.Context) (ignored bool) {
	if atomic.LoadInt32(&o.processing) == 0 {
		return
	}
	time.Sleep(time.Millisecond)
	return o.onAfter(ctx)
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

func (o *Exporter) onPanic(_ context.Context, _ interface{}) {
	cores.Registry.Debugger("%s panic: %v", o.processor.Name())
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: init and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) init() *Exporter {
	o.processor = process.New("term-tracer-exporter").
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall)
	return o
}
