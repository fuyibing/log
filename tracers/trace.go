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
	// 跨度组件.
	Trace interface {
		// Context
		// 上下文.
		Context() context.Context

		// Kv
		// 链路Key/Value属性.
		Kv() loggers.Kv

		// Name
		// 链路名.
		Name() string

		// New
		// 创建跨度.
		//
		// 基于 Trace 生成新的链路跨度 Span.
		New(name string) Span

		// SpanId
		// 跨度ID.
		SpanId() SpanId

		// TraceId
		// 链路ID.
		TraceId() TraceId
	}

	// 链路组件.
	trace struct {
		ctx  context.Context
		kv   loggers.Kv
		name string

		spanId  SpanId
		traceId TraceId
	}
)

// NewTrace
// 链路跟踪.
func NewTrace(name string) Trace {
	v := (&trace{name: name}).init()
	v.traceId = Operator.Generator().TraceIdNew()
	v.ctx = context.WithValue(context.Background(), ContextKey, v)
	return v
}

// NewTraceFromContext
// 链路跟踪.
func NewTraceFromContext(ctx context.Context, name string) Trace {
	// 复用上下文.
	if g := ctx.Value(ContextKey); g != nil {
		if v, ok := g.(Trace); ok {
			return v
		}
	}

	// 新建上下文.
	v := (&trace{name: name}).init()
	v.traceId = Operator.Generator().TraceIdNew()
	v.ctx = context.WithValue(ctx, ContextKey, v)
	return v
}

// NewTraceFromRequest
// 链路跟踪.
func NewTraceFromRequest(req *http.Request, name string) Trace {
	v := (&trace{name: name}).init()
	v.parseRequestField(req)

	// 解析ID.
	if !v.parseRequestId(req) {
		v.traceId = Operator.Generator().TraceIdNew()
		v.spanId = Operator.Generator().SpanIdNew()
	}

	// 配置上下文.
	v.ctx = context.WithValue(req.Context(), ContextKey, v)
	return v
}

func (o *trace) Context() context.Context { return o.ctx }
func (o *trace) Kv() loggers.Kv           { return o.kv }
func (o *trace) Name() string             { return o.name }

func (o *trace) New(name string) Span {
	v := (&span{name: name}).init()
	v.kv.Copy(o.kv)
	v.parentSpanId = o.spanId
	v.trace = o

	v.ctx = context.WithValue(o.ctx, ContextKey, v)
	return v
}

func (o *trace) SpanId() SpanId   { return o.spanId }
func (o *trace) TraceId() TraceId { return o.traceId }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *trace) init() *trace {
	o.kv = loggers.Kv{}
	return o
}

func (o *trace) parseRequestField(req *http.Request) {
	o.kv.Add(
		"http.protocol", req.Proto,
	).Add(
		"http.request.header", req.Header,
	).Add(
		"http.request.method", req.Method,
	).Add(
		"http.request.url", req.RequestURI,
	).Add(
		"http.user.agent", req.UserAgent(),
	)
}

func (o *trace) parseRequestId(req *http.Request) bool {
	// 获取 X-B3 Trace id.
	//
	//   {
	//     "X-B3-Traceid": "trace id"
	//   }
	if s := req.Header.Get(configurer.Config.GetOpenTracingTraceId()); s != "" {
		if v := Operator.Generator().TraceIdFromHex(s); v.IsValid() {
			o.traceId = v
		}
	}

	// 获取 X-B3 Span id.
	//
	//   {
	//     "X-B3-Spanid": "span id"
	//   }
	if s := req.Header.Get(configurer.Config.GetOpenTracingSpanId()); s != "" {
		if v := Operator.Generator().SpanIdFromHex(s); v.IsValid() {
			o.spanId = v
		}
	}

	return o.traceId.IsValid() && o.spanId.IsValid()
}
