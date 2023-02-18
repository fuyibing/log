// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package kafka

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formatters"
)

type (
	Executor struct {
		formatter formatters.Formatter
	}
)

func New() *Executor {
	return (&Executor{}).init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Executor) Logs(lines ...*base.Line) error { return nil }

func (o *Executor) SetFormatter(formatter formatters.Formatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Executor) init() *Executor {
	return o
}
