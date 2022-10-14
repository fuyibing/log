// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package adapters

import (
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

const (
	TracingCtx = "FybOpenTracingCtx"
	TracingKey = "FybOpenTracingKey"
)

type Tracing struct {
	spanId       string
	traceId      string
	offset       int32
	parentSpanId string
	prefix       string
}

func NewTracing() *Tracing {
	return &Tracing{}
}

func (o *Tracing) GetOffset() int32 {
	return atomic.LoadInt32(&o.offset)
}

func (o *Tracing) GetOffsetIncr() int32 {
	n := o.GetOffset()
	atomic.AddInt32(&o.offset, 1)
	return n
}

func (o *Tracing) GetParentSpanId() string { return o.parentSpanId }
func (o *Tracing) GetPrefix() string       { return o.prefix }
func (o *Tracing) GetSpanId() string       { return o.spanId }
func (o *Tracing) GetTraceId() string      { return o.traceId }

func (o *Tracing) WithRequest() {}

func (o *Tracing) WithParent(x *Tracing) *Tracing {
	o.traceId = x.GetTraceId()
	o.parentSpanId = o.GetSpanId()
	o.spanId = o.uuid()
	o.prefix = fmt.Sprintf("%s.%d", x.GetPrefix(), x.GetOffset())
	return o
}

func (o *Tracing) WithStart() *Tracing {
	o.traceId = o.uuid()
	o.parentSpanId = o.traceId
	o.spanId = o.traceId
	o.prefix = "0"
	return o
}

// 生成唯一ID.
func (o *Tracing) uuid() string {
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}
