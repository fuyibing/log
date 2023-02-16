// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package adapters

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formats"
)

type (
	AdapterInterface interface {
		Child(AdapterInterface)
		Send(line *base.Line)
		SetFormatter(formatter formats.Formatter)
	}
)
