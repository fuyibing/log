// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package formats

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	TermFormatter struct {
	}
)

func (o *TermFormatter) Format(line *base.Line) (string, error) {
	return "term: formatted", nil
}

func (o *TermFormatter) Init() *TermFormatter {
	return o
}
