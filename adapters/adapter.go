// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package adapters

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formatters"
)

type (
	AdapterIgnore func(err error, lines ...*base.Line)

	AdapterManager interface {
		SetFormatter(v formatters.Formatter)
		Log(lines ...*base.Line)
		SetChild(v AdapterManager)
		SetIgnore(v AdapterIgnore)
	}
)
