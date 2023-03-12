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
)

const (
	ContextKey = "__LOG_TRACE_CONTEXT_KEY__"
)

type (
	// Trace
	// component for tracer.
	Trace interface {
		Context() context.Context
		Kv() loggers.Kv
		Name() string
		New(name string) Span
		SpanId() SpanId
		TraceId() TraceId
	}

	trace struct {
		ctx  context.Context
		kv   loggers.Kv
		name string

		spanId  SpanId
		traceId TraceId
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *trace) Context() context.Context { return o.ctx }
func (o *trace) Kv() loggers.Kv           { return o.kv }
func (o *trace) Name() string             { return o.name }
func (o *trace) New(name string) Span     { return o.new(name) }
func (o *trace) SpanId() SpanId           { return o.spanId }
func (o *trace) TraceId() TraceId         { return o.traceId }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *trace) init() *trace {
	o.kv = loggers.Kv{}
	return o
}

func (o *trace) new(name string) Span {
	v := (&span{name: name}).init()
	v.kv.Copy(o.kv)
	v.parentSpanId = o.spanId
	v.trace = o

	v.ctx = context.WithValue(o.ctx, ContextKey, v)
	return v
}

func (o *trace) parseRequestField(req *http.Request) {
	o.kv.Add("http.protocol", req.Proto).
		Add("http.request.header", req.Header).
		Add("http.request.method", req.Method).
		Add("http.request.url", req.RequestURI).
		Add("http.user.agent", req.UserAgent())
}

func (o *trace) parseRequestId(req *http.Request) {
	// Read trace id.
	//
	//   {
	//     "X-B3-Traceid": "trace id"
	//   }
	if s := req.Header.Get(configurer.Config.GetOpenTracingTraceId()); s != "" {
		if v := Operator.Generator().TraceIdFromHex(s); v.IsValid() {
			o.traceId = v
		}
	}

	// Read Span id.
	//
	//   {
	//     "X-B3-Spanid": "span id"
	//   }
	if s := req.Header.Get(configurer.Config.GetOpenTracingSpanId()); s != "" {
		if v := Operator.Generator().SpanIdFromHex(s); v.IsValid() {
			o.spanId = v
		}
	}
}
