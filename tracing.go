// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

// Tracing interface.
type TracingInterface interface {
	Increment() int32
	ParentSpanId() string
	RequestMethod() string
	RequestUrl() string
	SpanId() string
	SpanVersion() string
	TraceId() string
	UseDefault() TracingInterface
	UseHeader(http.Header) TracingInterface
	UseRequest(*http.Request) TracingInterface
	Version(int32) string
}

// Tracing struct.
type tracing struct {
	parentSpanId  string
	spanId        string
	spanOffset    int32
	spanVersion   string
	traceId       string
	requestUrl    string
	requestMethod string
}

// New tracing instance.
func NewTracing() TracingInterface {
	o := new(tracing)
	o.spanId = o.generateUniqId()
	o.spanVersion = "0"
	return o
}

// Increment current offset and return old offset.
func (o *tracing) Increment() int32 {
	i := o.spanOffset
	atomic.AddInt32(&o.spanOffset, 1)
	return i
}

// Return parent span id.
func (o *tracing) ParentSpanId() string {
	return o.parentSpanId
}

func (o *tracing) RequestMethod() string { return o.requestMethod }
func (o *tracing) RequestUrl() string    { return o.requestUrl }

// Return span id.
func (o *tracing) SpanId() string {
	return o.spanId
}

// Return span version.
func (o *tracing) SpanVersion() string {
	return o.spanVersion
}

// Return trace id.
func (o *tracing) TraceId() string {
	return o.traceId
}

// Use default fields value.
func (o *tracing) UseDefault() TracingInterface {
	o.traceId = o.spanId
	return o
}

// Assign fields value by request header.
func (o *tracing) UseHeader(header http.Header) TracingInterface {
	// trace id.
	if traceId := header.Get(Config.NameTraceId); traceId != "" {
		o.traceId = traceId
	} else {
		o.traceId = o.spanId
	}
	// span id.
	if spanId := header.Get(Config.NameSpanId); spanId != "" {
		o.parentSpanId = spanId
	}
	// span version.
	if spanVersion := header.Get(Config.NameSpanVersion); spanVersion != "" {
		o.spanVersion = spanVersion
	}
	// end.
	return o
}

// Assign fields value by request.
func (o *tracing) UseRequest(req *http.Request) TracingInterface {
	o.UseHeader(req.Header)
	if method := req.Method; method != "" {
		o.requestMethod = method
	}
	if path := req.URL.Path; path != "" {
		o.requestUrl = path
	}
	return o
}

// Generate version.
func (o *tracing) Version(offset int32) string {
	return fmt.Sprintf("%s.%d", o.spanVersion, offset)
}

// Generate universally unique identifier.
func (o *tracing) generateUniqId() string {
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}
