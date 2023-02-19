// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package trace

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"net/http"
)

func NewContext() context.Context {
	t := (&Tracing{}).init()
	t.TraceId = t.genTraceId()

	// Return created with value.
	return context.WithValue(context.Background(),
		conf.OpenTracingKey,
		t,
	)
}

func NewChild(ctx context.Context) context.Context {
	if ctx != nil {
		if cv := ctx.Value(conf.OpenTracingKey); cv != nil {
			if ct, ok := cv.(*Tracing); ok {
				t := (&Tracing{}).init()
				t.WithParent(ct)
				t.ParentSpanId = ct.SpanId
				t.TraceId = ct.TraceId
				t.Version = ct.GenVersion(ct.GetPrevious())

				// Inherit parent http properties.

				if ct.Http {
					t.Http = true
					t.HttpHeaders = ct.HttpHeaders
					t.HttpProtocol = ct.HttpProtocol
					t.HttpRequestMethod = ct.HttpRequestMethod
					t.HttpRequestUrl = ct.HttpRequestUrl
					t.HttpUserAgent = ct.HttpUserAgent
				}

				return context.WithValue(context.Background(), conf.OpenTracingKey, t)
			}
		}
	}

	return NewContext()
}

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

	// Return created with value.
	return context.WithValue(context.Background(),
		conf.OpenTracingKey,
		t,
	)
}
