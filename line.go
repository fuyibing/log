// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

var (
	regexpLineDuration = regexp.MustCompile(`\[d=(\d+\.?\d*)\]`)
)

// Log Line.
type Line struct {
	duration float64
	level    Level
	text     string
	offset   int64
	time     time.Time
	tracing  *Tracing
}

// 创建日志行实例.
func NewLine(ctx interface{}, level Level, text string, args ...interface{}) *Line {
	line := &Line{time: time.Now(), level: level, text: fmt.Sprintf(text, args...)}
	line.parseCtx(ctx)
	line.parseDuration()
	return line
}

func (o *Line) GetLevel() Level     { return o.level }
func (o *Line) GetOffset() int64    { return o.offset }
func (o *Line) GetTimeline() string { return o.time.Format(Config.TimeFormat) }

func (o *Line) GetTracing() *Tracing {
	return o.tracing
}

func (o *Line) HasTracing() bool {
	return o.tracing != nil
}

func (o *Line) String() string { return o.text }

// 解析链路.
func (o *Line) parseCtx(ctx interface{}) {
	// nil.
	if ctx == nil {
		return
	}
	// Use iris.Context.
	if ir, ok := ctx.(iris.Context); ok {
		if x := ir.Values().Get(OpenTracingKey); x != nil {
			o.tracing = x.(*Tracing)
			o.offset, _ = o.tracing.Increment()
		}
		return
	}
	// Use context.Context.
	if cc, ok := ctx.(context.Context); ok {
		if x := cc.Value(OpenTracingKey); x != nil {
			o.tracing = x.(*Tracing)
			o.offset, _ = o.tracing.Increment()
		}
		return
	}
	// Use TraceInterface
	if ti, ok := ctx.(*Tracing); ok {
		o.tracing = ti
		o.offset, _ = o.tracing.Increment()
		return
	}
}

// 解析时长.
func (o *Line) parseDuration() {
	if m := regexpLineDuration.FindStringSubmatch(o.text); len(m) == 2 {
		if d, e := strconv.ParseFloat(m[1], 64); e == nil {
			o.duration = d
		}
	}
}
