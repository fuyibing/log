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
	"github.com/fuyibing/log/v5/conf"
	"net/http"
)

type (
	// Trace
	// 链路接口.
	Trace interface {
		// GetContext
		// 获取 Trace 上下文.
		GetContext() context.Context

		// GetTraceId
		// 获取 TraceId.
		GetTraceId() TraceId

		// NewSpan
		// 基于此 Trace 创建 Span 跨度.
		NewSpan(name string) Span
	}

	trace struct {
		attr    Attr
		ctx     context.Context
		name    string
		spanId  SpanId
		traceId TraceId
	}
)

// NewTrace
// returns a root cores.Trace component.
func NewTrace(name string) Trace {
	return NewTraceFromContext(
		context.Background(),
		name,
	)
}

// NewTraceFromContext
// returns a cores.Trace component with context values.
func NewTraceFromContext(ctx context.Context, name string) Trace {
	// Reuse previous Trace.
	if v := ctx.Value(base.ContentKeyTrace); v != nil {
		if vc, ok := v.(Trace); ok {
			return vc
		}
	}

	// Return a Trace component with default properties.
	o := (&trace{name: name}).
		init().
		initContext(ctx)

	o.traceId = Identify.GenTraceId()
	return o
}

// NewTraceFromRequest
// returns a Trace component from a http request.
func NewTraceFromRequest(req *http.Request, name string) Trace {
	o := (&trace{name: name}).
		init().
		initContext(req.Context()).
		initRequest(req)
	return o
}

// GetContext
// 获取 Trace 上下文.
func (o *trace) GetContext() context.Context {
	return o.ctx
}

// GetTraceId
// 获取 TraceId.
func (o *trace) GetTraceId() TraceId {
	return o.traceId
}

// NewSpan
// 基于此 Trace 创建 Span 跨度.
func (o *trace) NewSpan(name string) Span {
	s := (&span{name: name}).
		init().
		initRelations(o, o.traceId, o.spanId).
		initContext(o.ctx)

	s.attr.Copy(o.attr)
	return s
}

// init
// 构造链路.
func (o *trace) init() *trace {
	o.attr = NewAttr()
	return o
}

// initContext
// 初始化 Trace 上下文.
func (o *trace) initContext(ctx context.Context) *trace {
	o.ctx = context.WithValue(ctx, base.ContextKeySpan, o)
	return o
}

// initRequest
// 基于 HTTP 请求参数, 初始化链路.
func (o *trace) initRequest(req *http.Request) *trace {
	// 请求参数.
	// 将 HTTP 请求参数加到 Trace 属性.
	o.attr.Add(base.ResourceHttpRequestUrl, req.RequestURI).
		Add(base.ResourceHttpRequestMethod, req.RequestURI).
		Add(base.ResourceHttpHeader, req.Header).
		Add(base.ResourceHttpUserAgent, req.UserAgent()).
		Add(base.ResourceHttpProtocol, req.Proto)

	// OpenTracing.
	// 解析 OpenTracing 参数为 TraceId.
	if ht := req.Header.Get(conf.Config.GetOpenTracingTraceId()); ht != "" {
		o.traceId = Identify.HexTraceId(ht)

		// 上游 Span 解析 SpanId.
		if hs := req.Header.Get(conf.Config.GetOpenTracingSpanId()); hs != "" {
			o.spanId = Identify.HexSpanId(hs)
		}

		return o
	}

	// 生成新 TraceId.
	o.traceId = Identify.GenTraceId()
	return o
}
