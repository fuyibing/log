// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package formats

import (
	"github.com/fuyibing/log/v8/base"
)

type (
	KafkaFormatter struct {
	}
)

func (o *KafkaFormatter) Format(line *base.Line) (string, error) {
	return "kafka: formatted", nil
}

func (o *KafkaFormatter) Init() *KafkaFormatter {
	return o
}
