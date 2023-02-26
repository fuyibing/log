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

package cores

import (
	"context"
	"github.com/fuyibing/log/v5/base"
	"sync"
)

type (
	// Span
	// is the individual component of a trace.
	Span interface {
		// AddEvent
		// add SpanEvent on the Span component.
		AddEvent(events ...SpanEvent) Span

		// Child
		// returns a child Span component.
		Child(name string) Span

		// End
		// stop Span recorder.
		End()

		// GetAttr
		// return Attr component of the Span.
		GetAttr() Attr

		// GetContext
		// returns a context.Context of the Span.
		GetContext() context.Context

		// GetLogs
		// return SpanLogs component of the Span.
		GetLogs() SpanLogs

		// GetName
		// return name of the Span.
		GetName() string

		// GetSpanId
		// returns a SpanId identify of the Span.
		GetSpanId() SpanId

		// GetParentSpanId
		// returns a parent SpanId identify of the Span.
		GetParentSpanId() SpanId

		// GetSpanTime
		// returns a SpanTime component of the Span.
		GetSpanTime() SpanTime

		// GetTrace
		// returns a Trace component of the Span.
		GetTrace() Trace

		// GetTraceId
		// returns a TraceId identify of the Span.
		GetTraceId() TraceId

		// Logger
		// returns a SpanLogger component.
		Logger() *SpanLogger
	}

	span struct {
		sync.RWMutex

		attr                 Attr
		ctx                  context.Context
		events               []SpanEvent
		name                 string
		spanId, parentSpanId SpanId
		spanLogs             SpanLogs
		spanTime             SpanTime
		trace                Trace
		traceId              TraceId
	}
)

// AddEvent
// add SpanEvent on the Span component.
func (o *span) AddEvent(events ...SpanEvent) Span {
	o.Lock()
	defer o.Unlock()
	o.events = append(o.events, events...)
	return o
}

// Child
// returns a child Span component.
func (o *span) Child(name string) Span {
	return (&span{name: name}).
		init().
		initRelations(o.trace, o.traceId, o.spanId).
		initContext(o.ctx)
}

// End
// stop Span recorder.
func (o *span) End() {
	o.spanTime.End()

	// Push span to exporter if enabled.
	if Registry.TracerEnabled() {
		Registry.TracerExporter().Push(o)
	}

	// Call registered events.
	for _, event := range o.events {
		event.Do(o)
	}
}

// GetAttr
// return Attr component of the Span.
func (o *span) GetAttr() Attr {
	return o.attr
}

// GetContext
// returns a context.Context of the Span.
func (o *span) GetContext() context.Context {
	return o.ctx
}

// GetLogs
// return SpanLogs component of the Span.
func (o *span) GetLogs() SpanLogs {
	o.RLock()
	defer o.RUnlock()
	return o.spanLogs
}

// GetName
// return name of the Span.
func (o *span) GetName() string {
	return o.name
}

// GetSpanId
// returns a SpanId identify of the Span.
func (o *span) GetSpanId() SpanId {
	return o.spanId
}

// GetParentSpanId
// returns a parent SpanId identify of the Span.
func (o *span) GetParentSpanId() SpanId {
	return o.parentSpanId
}

// GetSpanTime
// returns a SpanTime component of the Span.
func (o *span) GetSpanTime() SpanTime {
	return o.spanTime
}

// GetTrace
// returns a Trace component of the Span.
func (o *span) GetTrace() Trace {
	return o.trace
}

// GetTraceId
// returns a TraceId identify of the Span.
func (o *span) GetTraceId() TraceId {
	return o.traceId
}

// Logger
// return a SpanLogger component.
func (o *span) Logger() *SpanLogger {
	return NewSpanLogger(o)
}

// init
// span fields initialize.
func (o *span) init() *span {
	o.attr = NewAttr()
	o.events = make([]SpanEvent, 0)
	o.spanId = Identify.GenSpanId()
	o.spanLogs = NewSpanLogs()
	o.spanTime = NewSpanTime()
	return o
}

// initContext
// initialize context.Context on Span.
func (o *span) initContext(ctx context.Context) *span {
	// Redirect background context
	// if nil set by param.
	if ctx == nil {
		ctx = context.Background()
	}

	// Bind a Span component
	// on context.Context with specified key.
	o.ctx = context.WithValue(ctx, base.ContextKeySpan, o)
	return o
}

// initRelations
// initialize TraceId, SpanId on Span.
func (o *span) initRelations(t Trace, tid TraceId, pid SpanId) *span {
	o.trace = t
	o.traceId = tid
	o.parentSpanId = pid
	return o
}
