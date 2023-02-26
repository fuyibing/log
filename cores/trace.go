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
	// is the creator of Span.
	Trace interface {
		GetContext() context.Context
		GetTraceId() TraceId
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
// returns a Trace component.
func NewTrace(name string) Trace {
	return NewTraceFromContext(
		context.Background(),
		name,
	)
}

// NewTraceFromContext
// returns a Trace component from a context.Context.
func NewTraceFromContext(ctx context.Context, name string) Trace {
	o := (&trace{name: name}).
		init().
		initContext(ctx).
		initTraceId()
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

func (o *trace) GetContext() context.Context {
	return o.ctx
}

func (o *trace) GetTraceId() TraceId {
	return o.traceId
}

func (o *trace) NewSpan(name string) Span {
	s := (&span{name: name}).
		init().
		initRelations(o, o.traceId, o.spanId).
		initContext(o.ctx)

	s.attr.Copy(o.attr)
	return s
}

func (o *trace) init() *trace {
	o.attr = NewAttr()
	return o
}

func (o *trace) initContext(ctx context.Context) *trace {
	o.ctx = context.WithValue(ctx, base.ContextKeySpan, o)
	return o
}

func (o *trace) initRequest(req *http.Request) *trace {
	// Append header options.
	o.attr.Add(base.ResourceHttpRequestUrl, req.RequestURI).
		Add(base.ResourceHttpRequestMethod, req.RequestURI).
		Add(base.ResourceHttpHeader, req.Header).
		Add(base.ResourceHttpUserAgent, req.UserAgent()).
		Add(base.ResourceHttpProtocol, req.Proto)

	// Parse Open Tracing values.
	if ht := req.Header.Get(conf.Config.GetOpenTracingTraceId()); ht != "" {
		// Hex trace id.
		o.traceId = Identify.HexTraceId(ht)

		// Hex span id.
		if hs := req.Header.Get(conf.Config.GetOpenTracingSpanId()); hs != "" {
			o.spanId = Identify.HexSpanId(hs)
		}

		return o
	}

	// Default trace id.
	o.traceId = Identify.GenTraceId()
	return o
}

func (o *trace) initTraceId() *trace {
	o.traceId = Identify.GenTraceId()
	return o
}
