// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package trace

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"net/http"
)

// New
// return root trace context. Compatible with old version.
func New() context.Context {
	return NewContext()
}

// NewContext
// return root trace context.
func NewContext() context.Context {
	// Prepare tracing
	// and generate trace id.
	t := (&Tracing{}).init()
	t.TraceId = t.genTraceId()

	// Root trace context
	// returned.
	return context.WithValue(root, conf.OpenTracingKey, t)
}

// NewChild
// return child trace context. Return root trace context if parent
// not specified.
func NewChild(ctx context.Context) context.Context {
	// Check parent tracing.
	if ctx != nil {
		if cv := ctx.Value(conf.OpenTracingKey); cv != nil {
			if ct, ok := cv.(*Tracing); ok {
				// Prepare tracing
				// and set required fields.
				t := (&Tracing{}).init().WithParent(ct)
				t.ParentSpanId = ct.SpanId
				t.TraceId = ct.TraceId
				t.Version = ct.GenVersion(ct.GetPrevious())

				// Inherit
				// parent fields.
				if ct.Http {
					t.Http = true
					t.HttpHeaders = ct.HttpHeaders
					t.HttpProtocol = ct.HttpProtocol
					t.HttpRequestMethod = ct.HttpRequestMethod
					t.HttpRequestUrl = ct.HttpRequestUrl
					t.HttpUserAgent = ct.HttpUserAgent
				}

				// Child trace context
				// returned.
				return context.WithValue(root, conf.OpenTracingKey, t)
			}
		}
	}

	// Use root trace
	// returned.
	return NewContext()
}

// NewRequest
// create trace context based on http request.
func NewRequest(req *http.Request) context.Context {
	t := (&Tracing{}).init()

	// Assign trace id.
	if s := req.Header.Get(conf.Config.GetTraceId()); s != "" {
		t.TraceId = s
	} else {
		t.TraceId = t.genTraceId()
	}

	// Assign parent span id.
	if s := req.Header.Get(conf.Config.GetSpanId()); s != "" {
		t.ParentSpanId = s
	}

	// Assign version.
	if s := req.Header.Get(conf.Config.GetTraceVersion()); s != "" {
		t.Version = s
	}

	// Assign request info.

	t.Http = true
	t.HttpHeaders = req.Header
	t.HttpProtocol = req.Proto
	t.HttpRequestUrl = req.RequestURI
	t.HttpRequestMethod = req.Method
	t.HttpUserAgent = req.UserAgent()

	// Request trace context
	// returned.
	return context.WithValue(root, conf.OpenTracingKey, t)
}
