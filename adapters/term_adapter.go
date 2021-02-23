// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"fmt"

	"github.com/fuyibing/log/v2/interfaces"
)

type termAdapter struct {
}

func (o *termAdapter) Run(line interfaces.LineInterface) {
	s := fmt.Sprintf("[%s][%s]", line.Timeline(), line.ColorLevel())
	if line.Tracing() {
		s += fmt.Sprintf("[v=%s]", line.SpanVersion())
	}
	s += " " + line.Content()
	fmt.Printf("%s\n", s)
}

func NewTerm() *termAdapter {
	return &termAdapter{}
}
