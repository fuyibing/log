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
// date: 2023-02-26

package log

import (
	"context"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/log/v5/exporters/logger_term"
	"github.com/fuyibing/log/v5/exporters/tracer_jaeger"
	"github.com/fuyibing/log/v5/exporters/tracer_term"
	"github.com/fuyibing/util/v8/process"
	"net/http"
	"sync"
	"time"
)

var (
	Manager Management
)

type (
	Management interface {
		NewTrace(name string) cores.Trace
		NewTraceFromContext(ctx context.Context, name string) cores.Trace
		NewTraceFromRequest(req *http.Request, name string) cores.Trace
		Start(ctx context.Context) error
		Stop()
	}

	management struct {
		processor process.Processor
	}
)

// NewTrace
// returns a root cores.Trace component.
func (o *management) NewTrace(name string) cores.Trace {
	return cores.NewTrace(name)
}

// NewTraceFromContext
// returns a cores.Trace component with context values.
func (o *management) NewTraceFromContext(ctx context.Context, name string) cores.Trace {
	return cores.NewTraceFromContext(ctx, name)
}

// NewTraceFromRequest
// returns a cores.Trace component with http request values.
func (o *management) NewTraceFromRequest(req *http.Request, name string) cores.Trace {
	return cores.NewTraceFromRequest(req, name)
}

// Start manager, block goroutine.
func (o *management) Start(ctx context.Context) error {
	return o.processor.Start(ctx)
}

// Stop manager.
func (o *management) Stop() {
	o.processor.Stop()

	// Sleep 100ms until processor stopped.
	for {
		if o.processor.Stopped() {
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Manager events
// /////////////////////////////////////////////////////////////////////////////

func (o *management) init() *management {
	o.processor = process.New("log-tracer").
		Before(o.onBeforeLogger, o.onBeforeTracer).
		Callback(o.onCallBefore, o.onCallListen).
		Panic(o.onPanic)
	return o
}

func (o *management) onBeforeLogger(_ context.Context) (ignored bool) {
	name := conf.Config.GetLoggerExporter()
	switch name {
	case "term":
		exporter := logger_term.NewExporter()
		o.processor.Add(exporter.Processor())
		cores.Registry.RegisterLoggerExporter(exporter)
	}
	return
}

func (o *management) onBeforeTracer(_ context.Context) (ignored bool) {
	name := conf.Config.GetTracerExporter()
	switch name {
	case "term":
		exporter := tracer_term.NewExporter()
		o.processor.Add(exporter.Processor())
		cores.Registry.RegisterTracerExporter(exporter)
	case "jaeger":
		exporter := tracer_jaeger.NewExporter()
		o.processor.Add(exporter.Processor())
		cores.Registry.RegisterTracerExporter(exporter)
	}
	return
}

func (o *management) onCallBefore(_ context.Context) (ignored bool) {
	cores.Registry.Update()
	return
}

func (o *management) onCallListen(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *management) onPanic(_ context.Context, v interface{}) {
	cores.Registry.Debugger("log.Manager fatal: %v", v)
}

func init() {
	new(sync.Once).Do(func() {
		Manager = (&management{}).init()
	})
}
