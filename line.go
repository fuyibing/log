// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/fuyibing/log/v2/interfaces"
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

// 日志行结构体.
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
	tracing     interfaces.TraceInterface
}

// 创建日志行实例.
func NewLine(ctx interface{}, level interfaces.Level, text string, args []interface{}) interfaces.LineInterface {
	// 行实例.
	o := &Line{
		time: time.Now(),
		text: text, args: args, level: level,
		pid:         Config.GetPid(),
		serviceName: Config.AppName(), serviceAddr: Config.AppAddr(),
	}
	// 执行时长.
	o.parseDuration()
	// 请求链.
	if tracer := ParseTracing(ctx); tracer != nil {
		o.tracing = tracer
		o.offset, _ = o.tracing.IncrOffset()
	}
	return o
}

// 返回带颜色Level文本.
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

// 日志正文.
func (o *Line) Content() string {
	if o.args != nil && len(o.args) > 0 {
		return fmt.Sprintf(o.text, o.args...)
	}
	return o.text
}

// 执行时长.
func (o *Line) Duration() float64 {
	return o.duration
}

// 日志级别.
func (o *Line) Level() string { return Config.GetLevel(o.level) }

// 上级Span.
func (o *Line) ParentSpanId() string {
	if o.tracing != nil {
		return o.tracing.GetParentSpanId()
	}
	return ""
}

// 进程ID.
func (o *Line) Pid() int {
	return o.pid
}

// 请求信息.
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
