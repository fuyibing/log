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

package tracer

import (
	"context"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"net/http"
)

type (
	// Trace 跟踪组件.
	Trace struct {
		// Attribute 属性定义.
		//
		// 跟踪组件 Key/Value 键值对, 转成JSON之后结果如下:
		//
		//   {
		//       "key": "value"
		//   }
		Attribute base.Attribute

		// Ctx 跟踪上下文.
		Ctx context.Context

		// Name 跟踪名称.
		Name string

		// SpanId 跨度ID.
		//
		// 长度为16个字符的跨度ID
		SpanId base.SpanId

		// TraceId 跟踪ID.
		//
		// 长度为32个字符的跟踪ID.
		TraceId base.TraceId
	}
)

// NewTrace 生成跟踪组件.
func NewTrace(name string) base.Trace {
	return NewTraceFromContext(
		context.Background(),
		name,
	)
}

// NewTraceFromContext 生成跟踪组件.
//
// 基于指定上下文(context.Context)生成跟踪组件.
func NewTraceFromContext(ctx context.Context, name string) base.Trace {
	// 组件复用.
	if v := ctx.Value(base.ContextKeyForTrace); v != nil {
		// 基于跨度Span.
		if vc, ok := v.(*Span); ok {
			return vc.GetTrace()
		}

		// 基于跟踪Trace.
		if vc, ok := v.(*Trace); ok {
			return vc
		}
	}

	// 新建组件.
	x := (&Trace{Name: name}).init()
	x.initTraceId()
	x.initContext(ctx)
	return x
}

// NewTraceFromRequest 生成跟踪组件.
//
// 基于HTTP请求生成跟踪组件.
func NewTraceFromRequest(req *http.Request, name string) base.Trace {
	x := (&Trace{Name: name}).init()
	x.initRequestArguments(req)
	x.initRequestId(req)
	x.initContext(req.Context())
	return x
}

// Context 获取上下文.
func (o *Trace) Context() context.Context { return o.Ctx }

// GetAttribute 获取组件属性.
func (o *Trace) GetAttribute() base.Attribute { return o.Attribute }

// GetName 获取跟踪名称.
func (o *Trace) GetName() string { return o.Name }

// GetSpanId 获取跨度ID.
func (o *Trace) GetSpanId() base.SpanId { return o.SpanId }

// GetTraceId 获取跟踪ID.
func (o *Trace) GetTraceId() base.TraceId { return o.TraceId }

// Span 生成跨度.
func (o *Trace) Span(name string) base.Span {
	x := (&Span{Name: name}).init()
	x.Trace = o
	x.ParentSpanId = o.SpanId
	x.Attribute.Copy(o.Attribute)

	x.initContext(o.Ctx)
	return x
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

// init
// 初始化跟踪.
func (o *Trace) init() *Trace {
	o.Attribute = base.Attribute{}
	return o
}

// initContext
// 初始化Context上下文.
func (o *Trace) initContext(ctx context.Context) {
	o.Ctx = context.WithValue(ctx, base.ContextKeyForTrace, o)
}

// initRequestArguments
// 初始化HTTP请求.
func (o *Trace) initRequestArguments(req *http.Request) {
	o.Attribute.Add(base.ResourceHttpHeader, req.Header)
	o.Attribute.Add(base.ResourceHttpProtocol, req.Proto)
	o.Attribute.Add(base.ResourceHttpRequestMethod, req.Method)
	o.Attribute.Add(base.ResourceHttpRequestUrl, req.RequestURI)
	o.Attribute.Add(base.ResourceHttpUserAgent, req.UserAgent())
}

// initRequestId
// 初始化HTTP请求.
func (o *Trace) initRequestId(req *http.Request) {
	// 解析HTTP头参数.
	if s1 := req.Header.Get(conf.Config.GetOpenTracingTraceId()); s1 != "" {
		// 跟踪ID.
		if o.TraceId = base.Id.TraceIdFromHex(s1); o.TraceId.IsValid() {
			// 跨度ID.
			if s2 := req.Header.Get(conf.Config.GetOpenTracingSpanId()); s2 != "" {
				if o.SpanId = base.Id.SpanIdFromHex(s2); o.SpanId.IsValid() {
					return
				}
			}
		}
	}

	// 创建默认参数.
	o.initTraceId()
	o.initSpanId()
}

// initSpanId
// 初始化跨度ID.
func (o *Trace) initSpanId() {
	o.SpanId = base.SpanId{}
}

// initTraceId
// 初始化跟踪ID.
func (o *Trace) initTraceId() {
	o.TraceId = base.Id.TraceIdNew()
}
