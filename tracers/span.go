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
// date: 2023-03-03

package tracers

import (
	"context"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"net/http"
	"sync"
	"time"
)

type (
	// Span
	// 链路跨度.
	Span interface {
		ApplyRequest(req *http.Request)
		Child(name string) Span
		Context() context.Context
		Duration() time.Duration
		End()
		Kv() loggers.Kv
		Logger() SpanLogger
		Logs() []loggers.Log
		Name() string
		ParentSpanId() SpanId
		SpanId() SpanId
		StartTime() time.Time
		Trace() Trace
	}

	span struct {
		sync.RWMutex

		kv   loggers.Kv
		ctx  context.Context
		name string

		logs                 []loggers.Log
		spanId, parentSpanId SpanId
		startTime, endTime   time.Time
		trace                Trace
	}
)

func (o *span) ApplyRequest(req *http.Request) {
	req.Header.Set(configurer.Config.GetOpenTracingTraceId(), o.trace.TraceId().String())
	req.Header.Set(configurer.Config.GetOpenTracingSpanId(), o.trace.SpanId().String())
	req.Header.Set(configurer.Config.GetOpenTracingSampled(), "1")
}

func (o *span) Child(name string) Span {
	v := (&span{name: name}).init()
	v.trace = o.trace
	v.parentSpanId = o.spanId
	v.ctx = context.WithValue(o.ctx, ContextKey, v)
	return v
}

func (o *span) Context() context.Context { return o.ctx }
func (o *span) Duration() time.Duration  { return o.endTime.Sub(o.startTime) }

func (o *span) End() {
	o.Lock()
	o.endTime = time.Now()
	o.Unlock()

	Operator.Push(o)
}

func (o *span) Kv() loggers.Kv     { return o.kv }
func (o *span) Logger() SpanLogger { return spanLoggerAcquire(o) }

func (o *span) Logs() []loggers.Log {
	o.RLock()
	defer o.RUnlock()

	return o.logs
}

func (o *span) Name() string         { return o.name }
func (o *span) ParentSpanId() SpanId { return o.parentSpanId }
func (o *span) SpanId() SpanId       { return o.spanId }
func (o *span) StartTime() time.Time { return o.startTime }
func (o *span) Trace() Trace         { return o.trace }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *span) addLog(log loggers.Log) {
	o.Lock()
	defer o.Unlock()

	o.logs = append(o.logs, log)
}

func (o *span) init() *span {
	o.kv = loggers.Kv{}
	o.logs = make([]loggers.Log, 0)
	o.spanId = Operator.Generator().SpanIdNew()
	o.startTime = time.Now()
	return o
}
