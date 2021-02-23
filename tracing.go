// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	"github.com/fuyibing/log/v2/interfaces"
)

// 请求链结构体.
// Open tracing struct.
type tracing struct {
	method       string
	offset       int32
	parentSpanId string
	spanId       string
	spanVersion  string
	traceId      string
	url          string
}

// 创建OpenTracing.
func NewTracing() interfaces.TraceInterface  { return &tracing{spanVersion: "0"} }
func (o *tracing) GenVersion(i int32) string { return fmt.Sprintf("%s.%d", o.spanVersion, i) }
func (o *tracing) GetSpanId() string         { return o.spanId }
func (o *tracing) GetSpanVersion() string    { return o.spanVersion }
func (o *tracing) GetTraceId() string        { return o.traceId }

// Offset加1, 返回加之前的Offset.
func (o *tracing) IncrOffset() (before int32, after int32) {
	before = atomic.LoadInt32(&o.offset)
	after = atomic.AddInt32(&o.offset, 1)
	return
}

// Http请求参数.
func (o *tracing) RequestInfo() (method string, url string) { return o.method, o.url }

// 使用默认模式.
func (o *tracing) UseDefault() interfaces.TraceInterface {
	o.offset = 0
	o.spanId = o.generateUniqId()
	o.spanVersion = "0"
	o.traceId = o.spanId
	return o
}

// 使用Request模式.
func (o *tracing) UseRequest(req *http.Request) interfaces.TraceInterface {
	o.offset = 0
	o.spanId = o.generateUniqId()
	o.method = req.Method
	o.url = req.URL.Path
	return o.parseHeader(req.Header)
}

// Generate universally unique identifier.
func (o *tracing) generateUniqId() string {
	if u, e := uuid.NewUUID(); e == nil {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	t := time.Now()
	return fmt.Sprintf("a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
}

// 使用默认模式.
func (o *tracing) parseHeader(header http.Header) interfaces.TraceInterface {
	ti, si, sv := Config.GetTrace()
	// Trace id.
	if x, ok := header[ti]; ok && len(x) > 0 && x[0] != "" {
		o.traceId = x[0]
	}
	if o.traceId == "" {
		o.traceId = o.spanId
	}
	// Span id.
	if x, ok := header[si]; ok && len(x) > 0 && x[0] != "" {
		o.parentSpanId = x[0]
	}
	// Span version
	if x, ok := header[sv]; ok && len(x) > 0 && x[0] != "" {
		o.spanVersion = x[0]
	}
	// with header.
	return o
}
