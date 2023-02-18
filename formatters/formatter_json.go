// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	JsonFormatter struct {
	}
)

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) Body(line *base.Line) []byte { return nil }

func (o *JsonFormatter) String(line *base.Line) string { return "" }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) init() *JsonFormatter {
	return o
}
