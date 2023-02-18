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

		offset, previous int32
		parent           *Tracing
	}
)

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Tracing) GenVersion(n int32) string {
	return fmt.Sprintf("%v.%d", o.Version, n)
}

func (o *Tracing) GetIncrement() int32 {
	o.previous = atomic.AddInt32(&o.offset, 1) - 1
	return o.previous
}

func (o *Tracing) GetPrevious() int32 {
	return o.previous
}

func (o *Tracing) WithParent(p *Tracing) {
	o.parent = p
}

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
