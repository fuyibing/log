// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	ErrorFormatter struct{}
)

func NewErrorFormatter() *ErrorFormatter {
	return &ErrorFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *ErrorFormatter) Body(line *base.Line) []byte {
	return []byte(o.String(line))
}

func (o *ErrorFormatter) String(line *base.Line) (str string) {
	str += "{ERR} " + line.Text
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *ErrorFormatter) init() *ErrorFormatter {
	return o
}
