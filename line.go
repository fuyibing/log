// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

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
	regexpLineDuration = regexp.MustCompile(`\[d=([^\]]+)\]`)
)

// Log line.
type Line struct {
	content       string
	duration      float64
	level         Level
	offset        int
	time          time.Time
	tracing       TracingInterface
	tracingOffset int32
}

// Log line interface.
type LineInterface interface {
	Duration() float64
	GetLevel() Level
	GetLevelText() string
	String() string
	Time() time.Time
	Timeline() string
	Tracing() TracingInterface
	TracingOffset() int32
}

// Create log line instance.
// ctx accept: context.Context, iris.Context.
func NewLine(ctx interface{}, level Level, text string, args ...interface{}) LineInterface {
	line := &Line{level: level, time: time.Now()}
	line.parseString(text, args)
	line.parseTracing(ctx)
	return line
}

// Return duration value.
func (o *Line) Duration() float64 {
	return o.duration
}

// Return log level.
func (o *Line) GetLevel() Level {
	return o.level
}

// Return log level name.
func (o *Line) GetLevelText() string {
	return LevelText[o.level]
}

// Return log content.
func (o *Line) String() string {
	return o.content
}

// Return time.
func (o *Line) Time() time.Time {
	return o.time
}

// Return log timeline.
func (o *Line) Timeline() string {
	return o.time.Format(Config.TimeFormat)
}

// Return tracing interface.
func (o *Line) Tracing() TracingInterface {
	return o.tracing
}

// Return tracing offset.
func (o *Line) TracingOffset() int32 {
	return o.tracingOffset
}

// Parse log content.
func (o *Line) parseString(text string, args []interface{}) {
	// format content.
	if args != nil && len(args) > 0 {
		o.content = fmt.Sprintf(text, args...)
	} else {
		o.content = text
	}
	// match duration.
	if m := regexpLineDuration.FindStringSubmatch(o.content); len(m) == 2 {
		if d, e := strconv.ParseFloat(m[1], 64); e == nil {
			o.duration = d
			o.content = regexpLineDuration.ReplaceAllString(o.content, "")
		}
	}
}

// Parse tracing from context.Context.
// ctx accept: context.Context, iris.Context
func (o *Line) parseTracing(ctx interface{}) {
	// empty context
	if ctx == nil {
		return
	}
	// context.Context
	if c, ok := ctx.(context.Context); ok && c != nil {
		o.parseTracingFromContext(c)
		return
	}
	// iris.Context
	if c, co := ctx.(iris.Context); co && c != nil {
		if v := c.Values().Get(OpenTracingContext); v != nil {
			o.parseTracingFromContext(v.(context.Context))
			return
		}
	}
}

// Parse tracing from context.Context.
func (o *Line) parseTracingFromContext(ctx context.Context) {
	c := ctx.Value(OpenTracingContext)
	if c == nil {
		return
	}
	if t, ok := c.(TracingInterface); ok {
		o.tracing = t
		o.tracingOffset = t.Increment()
	}
}
