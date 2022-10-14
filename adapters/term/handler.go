// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package term

import (
	"fmt"
	"os"

	"github.com/fuyibing/log/v3/adapters"
)

type (
	// Handler
	// 终端适配器.
	Handler interface {
		// Run
		// 提交日志.
		Run(line *adapters.Line, err error)

		// Start
		// 启动服务.
		Start()
	}

	handler struct{}
)

var (
	color = map[adapters.Level][]int{
		adapters.Debug: {30, 0}, // 白底黑字
		adapters.Info:  {32, 0}, // 白底绿字
		adapters.Warn:  {33, 0}, // 白底黄字
		adapters.Error: {31, 0}, // 白底红字
	}
)

func New() Handler {
	return (&handler{}).init()
}

// Run
// 打印日志.
func (o *handler) Run(line *adapters.Line, err error) {
	if line == nil {
		if err != nil {
			o.output(err.Error())
		}
		return
	}

	// 1. 监听结束.
	defer func() {
		if v := recover(); v != nil {
			o.output(fmt.Sprintf("term panic: %v", v))
		}

		// println("line: ", line.GetId(), " -> ", line.GetAcquires())
		line.Release()
	}()

	// 2. 基础信息.
	var (
		format = line.TimeFormat
		level  = line.Level.String()
		text   = ""
	)
	if Config.TimeFormat != "" {
		format = Config.TimeFormat
	}
	text += fmt.Sprintf("[%s][%s]", line.Time.Format(format), level)

	// 3. 链路信息.
	if line.Trace {
		text += fmt.Sprintf("[%s=%s.%d]", line.SpanId, line.SpanPrefix, line.SpanOffset)
	}

	// 4. 输出内容.
	text += fmt.Sprintf(" %s", line.Content)
	if err != nil {
		text += fmt.Sprintf(" << interrupted: %v", err)
	}

	// 5. 内容着色.
	if Config.Color {
		if Config.Color {
			if c, ok := color[line.Level]; ok {
				text = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, c[1], c[0], text, 0x1B)
			}
		}
	}

	o.output(text)
}

// Start
// 启动服务.
func (o *handler) Start() {}

// 构造实例.
func (o *handler) init() *handler { return o }

// 输出内容.
func (o *handler) output(text string) {
	if _, err := fmt.Fprintf(os.Stdout, "%s\n", text); err != nil {
		fmt.Printf("%s\n", text)
	}
}
