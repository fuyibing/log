// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package formatters

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	Formatter func(line *base.Line) string
)
