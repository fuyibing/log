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
func (o *handler) Log(line *base.Line, err error) {
	if line == nil {
		if err != nil {
			o.output(err.Error())
		}
		return
	}

	defer func() {
		line.Release()
		recover()
	}()

	text := formatters.Formatter.AsTerm(line, err)
	if *Config.Color {
		if c, ok := Colors[line.Level]; ok {
			text = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, c[1], c[0], text, 0x1B)
		}
	}
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

func (o *handler) output(text string) {
	if _, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", text)); err != nil {
		println(text)
	}
}