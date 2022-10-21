// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package formatters

import (
	"github.com/fuyibing/log/v3/base"
)

// JSON格式化.
// 格式化后的数据用于写入到Redis/Kafka中.
type jsonFormatter struct{}

// Format
// 格式化过程.
func (o *jsonFormatter) Format(line *base.Line) string {
	return NewData(line).String()
}

// 构造实例.
func (o *jsonFormatter) init() *jsonFormatter {
	return o
}
