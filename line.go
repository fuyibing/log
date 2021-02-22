// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/fuyibing/log/interfaces"
)

var (
	colors = map[interfaces.Level][]int{
		interfaces.LevelDebug: {30, 47},
		interfaces.LevelInfo:  {37, 44},
		interfaces.LevelWarn:  {31, 43},
		interfaces.LevelError: {33, 41},
	}
	regexpLineDuration = regexp.MustCompile(`\[d=(\d+\.?\d*)\]`)
)

type Line struct {
	args        []interface{}
	duration    float64
	text        string
	level       interfaces.Level
	offset      int32
	pid         int
	serviceName string
	serviceAddr string
	time        time.Time
	tracing     *tracing
}

func NewLine(ctx context.Context, level interfaces.Level, text string, args []interface{}) interfaces.LineInterface {
	o := &Line{
		time: time.Now(),
		text: text, args: args, level: level,
		pid:         Config.GetPid(),
		serviceName: Config.AppName(), serviceAddr: Config.AppAddr(),
	}
	o.parseDuration()
	o.parseTracing(ctx)
	return o
}

func (o *Line) ColorLevel() string {
	if c, ok := colors[o.level]; ok {
		return fmt.Sprintf("%c[%d;%d;%dm%5s%c[0m",
			0x1B, 0,
			c[1], c[0],
			o.Level(),
			0x1B,
		)
	}
	return o.Level()
}

func (o *Line) Content() string {
	if o.args != nil && len(o.args) > 0 {
		return fmt.Sprintf(o.text, o.args...)
	}
	return o.text
}

func (o *Line) Duration() float64 {
	return o.duration
}

func (o *Line) Level() string { return Config.GetLevel(o.level) }

func (o *Line) ParentSpanId() string {
	if o.tracing != nil {
		return o.tracing.parentSpanId
	}
	return ""
}

func (o *Line) Pid() int {
	return o.pid
}

func (o *Line) RequestInfo() (method string, url string) {
	if o.tracing != nil {
		method, url = o.tracing.RequestInfo()
	}
	return
}

func (o *Line) ServiceAddr() string { return o.serviceAddr }

func (o *Line) ServiceName() string { return o.serviceName }

func (o *Line) SpanId() string {
	if o.tracing != nil {
		return o.tracing.GetSpanId()
	}
	return ""
}

func (o *Line) SpanVersion() string {
	if o.tracing != nil {
		return o.tracing.GenVersion(o.offset)
	}
	return ""
}

func (o *Line) Time() time.Time {
	return o.time
}

func (o *Line) Timeline() string {
	return o.time.Format(Config.GetTimeFormat())
}

func (o *Line) TraceId() string {
	if o.tracing != nil {
		return o.tracing.GetTraceId()
	}
	return ""
}

func (o *Line) Tracing() bool {
	return o.tracing != nil
}

// Parse duration from text.
func (o *Line) parseDuration() {
	if m := regexpLineDuration.FindStringSubmatch(o.text); len(m) == 2 {
		if d, e := strconv.ParseFloat(m[1], 64); e == nil {
			o.duration = d
		}
	}
}

// Parse tracing from Context.
func (o *Line) parseTracing(ctx context.Context) {
	if ctx == nil {
		return
	}
	if get := ctx.Value(interfaces.OpenTracingKey); get != nil {
		if tracing, ok := get.(*tracing); ok {
			o.tracing = tracing
			o.offset, _ = o.tracing.IncrOffset()
		}
	}
}
