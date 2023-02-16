// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package file

import (
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formats"
)

type (
	Executor struct {
		child     adapters.AdapterInterface
		formatter formats.Formatter
	}
)

func (o *Executor) Child(adapter adapters.AdapterInterface) { o.child = adapter }

func (o *Executor) Send(line *base.Line) {
	println("file: send")

	if o.child != nil {
		o.child.Send(line)
	}
}

func (o *Executor) SetFormatter(formatter formats.Formatter) { o.formatter = formatter }

func (o *Executor) Init() *Executor {
	o.formatter = (&formats.FileFormatter{}).Init().Format
	return o
}
