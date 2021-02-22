// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"github.com/fuyibing/log/interfaces"
)

type fileAdapter struct {
}

func (o *fileAdapter) Run(lineInterface interfaces.LineInterface) {}

func NewFile() *fileAdapter {
	return &fileAdapter{}
}

