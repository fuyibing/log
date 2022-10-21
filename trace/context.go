// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package trace

import (
	"context"
	"fmt"
	"net/http"
)

// New
// 创建链路.
//
// 返回链路控制上下文, 具有相同上下文的日志会被记录上链路属性.
func New() context.Context {
	return context.WithValue(context.Background(),
		OpenTracingKey, NewTracing().WithRoot(),
	)
}

// Child
// 创建子链路.
func Child(ctx context.Context, text string, args ...interface{}) context.Context {
	if tracing, ok := ctx.Value(OpenTracingKey).(*Tracing); ok {
		return context.WithValue(ctx, OpenTracingKey, NewTracing().WithTracing(tracing))
	}
	return New()
}

// FromRequest
// 从HTTP请求中解析.
func FromRequest(request *http.Request) context.Context {
	if s := request.Header.Get(TracingTraceId); s != "" {
		return context.WithValue(context.Background(), OpenTracingKey, NewTracing().WithRequest(request))
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
