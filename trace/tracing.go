// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package trace

import (
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/google/uuid"
)

// Tracing
// 链路参数.
type Tracing struct {
	TraceId      string // 主链路
	ParentSpanId string // 上级链路
	SpanId       string // 本级链路
	SpanPrefix   string // 链路前缀
	SpanOffset   int32  // 链路偏移

	RequestMethod string
	RequestUrl    string
}

// NewTracing
// 创建链路.
func NewTracing() *Tracing {
	o := &Tracing{}
	return o
}

// Increment
// 偏移量递加.
func (o *Tracing) Increment() int32 {
	return atomic.AddInt32(&o.SpanOffset, 1)
}

// WithRequest
// 基于HTTP请求.
func (o *Tracing) WithRequest(request *http.Request) *Tracing {
	o.TraceId = request.Header.Get(TracingTraceId)
	o.ParentSpanId = request.Header.Get(TracingSpanId)
	o.SpanId = o.makeSpanId()
	o.SpanPrefix = request.Header.Get(TracingSpanVersion)

	o.RequestMethod = request.Method
	o.RequestUrl = request.URL.RequestURI()
	return o
}

// WithRoot
// 基础根链路.
func (o *Tracing) WithRoot() *Tracing {
	o.SpanId = o.makeSpanId()
	o.SpanPrefix = defaultSpanPrefix
	o.TraceId = o.makeTraceId()
	return o
}

// WithTracing
// 基于链路.
func (o *Tracing) WithTracing(tracing *Tracing) *Tracing {
	o.TraceId = tracing.TraceId
	o.ParentSpanId = tracing.SpanId
	o.SpanId = o.makeSpanId()
	o.SpanPrefix = fmt.Sprintf("%s.%d", tracing.SpanPrefix, tracing.SpanOffset)
	return o
}

// 生成级别链路.
func (o *Tracing) makeSpanId() string {
	if u, err := uuid.NewUUID(); err == nil {
		if s := u.String(); len(s) >= 32 {
			return strings.ReplaceAll(u.String(), "-", "")[0:16]
		}
	}
	return o.makeTraceId()
}

// 生成主链路.
func (o *Tracing) makeTraceId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
