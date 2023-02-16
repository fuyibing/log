// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package formats

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	FileFormatter struct {
	}
)

func (o *FileFormatter) Format(line *base.Line) (string, error) {
	return "file: formatted", nil
}

func (o *FileFormatter) Init() *FileFormatter {
	return o
}
