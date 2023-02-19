// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package base

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/trace"
	"time"
)

type (
	// Line
	// log line definitions.
	Line struct {
		Duration float64
		Level    conf.Level
		Text     string
		Time     time.Time

		ctx           context.Context
		tracing       *trace.Tracing
		tracingOffset int32
	}
)

// Release
// instance to pool.
func (o *Line) Release() {
	Pool.ReleaseLine(o)
}

// TextParse
// collect variables into fields from text.
func (o *Line) TextParse() { Parser.Parse(o) }

// Tracing
// return tracing instance if configured.
func (o *Line) Tracing() *trace.Tracing {
	return o.tracing
}

// TracingOffset
// return line offset of trace span.
func (o *Line) TracingOffset() int32 {
	return o.tracingOffset
}

// WithContext
// read tracing from context, increment offset if success
// checked.
func (o *Line) WithContext(ctx context.Context) *Line {
	if o.ctx = ctx; o.ctx != nil {
		if cv := o.ctx.Value(conf.OpenTracingKey); cv != nil {
			if tracing, ok := cv.(*trace.Tracing); ok {
				o.tracingOffset = tracing.GetIncrement()
				o.tracing = tracing
			}
		}
	}
	return o
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Line) after() *Line {
	o.ctx = nil
	o.tracing = nil
	o.tracingOffset = 0

	o.Duration = 0
	o.Level = ""
	o.Text = ""
	return o
}

func (o *Line) before() *Line {
	o.Time = time.Now()
	return o
}

func (o *Line) init() *Line {
	return o
}
