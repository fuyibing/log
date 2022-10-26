// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package trace

import (
	"context"
	"fmt"
	"net/http"
)

// New
// 创建根上下文.
func New() context.Context {
	return NewContext()
}

func NewContext() context.Context {
	return context.WithValue(context.Background(),
		OpenTracingKey,
		NewTracing().WithRoot(),
	)
}

// Child
// 创建子上下文.
func Child(ctx context.Context) context.Context {
	if tracing, ok := ctx.Value(OpenTracingKey).(*Tracing); ok {
		return context.WithValue(ctx,
			OpenTracingKey,
			NewTracing().WithTracing(tracing),
		)
	}
	return New()
}

// BindRequest
// 绑定HTTP请求.
func BindRequest(ctx context.Context, request *http.Request) {
	if tracing, ok := ctx.Value(OpenTracingKey).(*Tracing); ok {
		request.Header.Set(TracingTraceId, tracing.TraceId)
		request.Header.Set(TracingParentSpanId, tracing.ParentSpanId)
		request.Header.Set(TracingSpanId, tracing.SpanId)
		request.Header.Set(TracingSpanVersion, fmt.Sprintf("%s.%d", tracing.SpanPrefix, tracing.SpanOffset))
	}
}

func FromContext(ctx context.Context) *Tracing {
	if v := ctx.Value(OpenTracingKey); v != nil {
		if t, ok := v.(*Tracing); ok {
			return t
		}
	}
	return nil
}

// FromRequest
// 从HTTP请求中解析.
func FromRequest(request *http.Request) context.Context {
	if s := request.Header.Get(TracingTraceId); s != "" {
		return context.WithValue(context.Background(),
			OpenTracingKey,
			NewTracing().WithRequest(request),
		)
	}
	return New()
}
