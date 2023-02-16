// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package formats

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	Formatter func(line *base.Line) (text string, err error)
)
