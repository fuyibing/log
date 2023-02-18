// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package errors

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formatters"
)

type (
	Executor struct{}
)

func New() *Executor { return (&Executor{}).init() }

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Executor) Logs(lines ...*base.Line) (err error) {
	for _, line := range lines {
		println("~", line.Text)
	}
	return
}

func (o *Executor) SetFormatter(_ formatters.Formatter) {}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Executor) init() *Executor { return o }
