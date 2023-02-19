// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/trace"
	"net/http"
)

// AddTrace
// read current trace then add to request.
func AddTrace(ctx context.Context, request *http.Request) {
	if ctx != nil {
		if cv := ctx.Value(conf.OpenTracingKey); cv != nil {
			if t, ok := cv.(*trace.Tracing); ok {
				request.Header.Set(conf.Config.GetTraceId(), t.TraceId)
				request.Header.Set(conf.Config.GetSpanId(), t.SpanId)
				request.Header.Set(conf.Config.GetTraceVersion(), t.GenVersion(t.GetPrevious()))
			}
		}
	}
}

// NewChild
// return new trace context, child of parent.
func NewChild(ctx context.Context) context.Context {
	return trace.NewChild(ctx)
}

// NewChildInfo
// set info level log before return new trace context, child of
// parent.
func NewChildInfo(ctx context.Context, text string, args ...interface{}) context.Context {
	c := trace.NewChild(ctx)
	Infofc(c, text, args...)
	return c
}

// NewContext
// return new root trace context.
func NewContext() context.Context {
	return trace.NewContext()
}

// NewContextInfo
// set info level log before return new root context.
func NewContextInfo(text string, args ...interface{}) context.Context {
	ctx := trace.NewContext()
	Infofc(ctx, text, args...)
	return ctx
}

// NewRequest
// create context based on http request. Return root context if
// parent trace not specified.
func NewRequest(request *http.Request) context.Context {
	return trace.NewRequest(request)
}

// NewRequestInfo
// set info level log before return. Create context based on http
// request. Return root context if parent trace not specified.
func NewRequestInfo(request *http.Request, text string, args ...interface{}) context.Context {
	c := trace.NewRequest(request)
	Infofc(c, text, args...)
	return c
}
