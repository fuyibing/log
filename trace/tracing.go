// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package trace

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/google/uuid"
)

type (
	// Tracing
	// mark open tracing.
	Tracing struct {
		ParentSpanId string
		SpanId       string
		TraceId      string
		Version      string

		Http                              bool
		HttpHeaders                       map[string][]string
		HttpProtocol                      string
		HttpRequestMethod, HttpRequestUrl string
		HttpUserAgent                     string

		offset, previous int32
		parent           *Tracing
	}
)

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

// GenVersion
// return version number of trace.
//
//   return "0.0"		// First log of the same tracer
//   return "0.3.0"
//   return "0.3.1"
func (o *Tracing) GenVersion(n int32) string { return fmt.Sprintf("%v.%d", o.Version, n) }

// GetIncrement
// increment offset then return previous value.
func (o *Tracing) GetIncrement() int32 {
	o.previous = atomic.AddInt32(&o.offset, 1) - 1
	return o.previous
}

// GetOffset
// return current offset.
func (o *Tracing) GetOffset() int32 { return o.offset }

// GetPrevious
// return previous offset value.
func (o *Tracing) GetPrevious() int32 { return o.previous }

// WithParent
// bind parent tracing to current.
func (o *Tracing) WithParent(p *Tracing) { o.parent = p }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Tracing) genSpanId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")[8:16]
}

func (o *Tracing) genTraceId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func (o *Tracing) init() *Tracing {
	o.SpanId = o.genSpanId()
	o.Version = "0"
	return o
}
