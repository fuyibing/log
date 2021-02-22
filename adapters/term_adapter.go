// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"github.com/fuyibing/log/interfaces"
)

type termAdapter struct {
}

func (o *termAdapter) Run(lineInterface interfaces.LineInterface) {}

func NewTerm() *termAdapter {
	return &termAdapter{}
}
