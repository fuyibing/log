// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package term

import (
	"context"
	"fmt"
	"os"

	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
)

type (
	handler struct{}
)

// New
// 创建执行器.
func New() base.AdapterEngine {
	return (&handler{}).init()
}

// Log
// 打印日志.
func (o *handler) Log(line *base.Line) {
	// 1. 无效日志.
	if line == nil {
		return
	}

	// 2. 监听结束.
	defer func() {
		// 2.1 实例回池.
		line.Release()

		// 2.2 捕获异常.
		if r := recover(); r != nil {
			o.output(fmt.Sprintf("panic on terminal adapter: %v", r))
		}
	}()

	// 3. 日志内容.
	text := formatters.Formatter.AsTerm(line)

	// 3.1 日志着色.
	if *Config.Color {
		if c, ok := Colors[line.Level]; ok {
			text = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, c[1], c[0], text, 0x1B)
		}
	}

	// 3.2 打印日志.
	o.output(text)
}

// Parent
// 绑定上级/降级.
func (o *handler) Parent(_ base.AdapterEngine) base.AdapterEngine {
	return o
}

// Start
// 启动服务.
func (o *handler) Start(_ context.Context) {
}

// 构造实例.
func (o *handler) init() *handler {
	return o
}

// 打印日志.
func (o *handler) output(text string) {
	if _, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", text)); err != nil {
		println(text)
	}
}
