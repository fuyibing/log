// author: wsfuyibing <websearch@163.com>
// date: 2021-03-15

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

// 调用链结构.
type Tracing struct {
	parentSpanId string
	spanId       string
	spanVersion  string
	traceId      string
	offset       int64
	method       string
	route        string
}

// ID递加.
func (o *Tracing) Increment() (cur int64, next int64) {
	cur = o.offset
	next = atomic.AddInt64(&o.offset, 1)
	return
}
func (o *Tracing) GetOffset() int64                       { return o.offset }
func (o *Tracing) GetParentSpanId() string                { return o.parentSpanId }
func (o *Tracing) GetRequestInfo() (method, route string) { return o.method, o.route }
func (o *Tracing) GetSpanId() string                      { return o.spanId }
func (o *Tracing) GetTraceId() string                     { return o.traceId }

// 默认模式.
func (o *Tracing) UseDefault() *Tracing {
	o.traceId = o.spanId
	return o
}

// 基于Request.
func (o *Tracing) UseRequest(req *http.Request) *Tracing {
	traceId, parentSpanId, spanVersion := Config.GetTrace()
	// Trace ID.
	//   header[X-B3-Traceid]
	if s := req.Header.Get(traceId); s != "" {
		o.traceId = s
	} else {
		o.traceId = o.spanId
	}
	// Parent span id.
	//   header[X-B3-Spanid]
	if s := req.Header.Get(parentSpanId); s != "" {
		o.parentSpanId = s
	}
	// Span version.
	//   header[X-B3-Version]
	if s := req.Header.Get(spanVersion); s != "" {
		o.spanVersion = s
	}
	// Request info.
	o.method = req.Method
	o.route = req.RequestURI
	return o
}

// 生成唯一ID.
func (o *Tracing) Uuid() string {
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}

// 格式化Version号.
func (o *Tracing) Version(n int64) string {
	return fmt.Sprintf("%s.%d", o.spanVersion, n)
}

func (o *Tracing) LinkVersion() string {
	return o.Version(o.offset - 1)
}

// 创建Tracing实例.
func NewTracing() *Tracing {
	o := &Tracing{offset: 0, spanVersion: "0"}
	o.spanId = o.Uuid()
	return o
}

func AssignRequest(req *http.Request, tracing *Tracing) {
	if tracing != nil {
		traceId, spanId, spanVersion := Config.GetTrace()
		req.Header.Set(traceId, tracing.GetTraceId())
		req.Header.Set(spanId, tracing.GetSpanId())
		req.Header.Set(spanVersion, tracing.Version(tracing.offset))
	}
}
