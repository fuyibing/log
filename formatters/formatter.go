// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	Formatter interface {
		Body(line *base.Line) []byte
		String(line *base.Line) string
	}
)
