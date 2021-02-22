// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"

	"github.com/fuyibing/log/interfaces"
)

type line struct {
}

func NewLine(ctx context.Context, level interfaces.Level, text string, args []interface{}) interfaces.LineInterface {
	return &line{}
}
