// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package formatters

import (
	"fmt"

	"github.com/fuyibing/log/v3/base"
)

// Term格式化.
// 格式化后的数据用于在终端上打印.
type termFormatter struct{}

// Format
// 格式化过程.
func (o *termFormatter) Format(line *base.Line) (text string) {
	// 1. 基础信息.
	text = fmt.Sprintf("[%s][%s]", line.Time.Format(base.LogTimeFormat), line.Level.Name())

	// 2. 链路信息.
	if line.Trace {
		text += fmt.Sprintf("[%s=%s.%d]", line.SpanId, line.SpanPrefix, line.SpanOffset)
	}

	// 3. 扩展信息.
	if line.Duration > 0 {
		text += fmt.Sprintf("[DURATION=%v]", line.Duration)
	}

	// 4. 日志内容.
	text += fmt.Sprintf(" %s", line.Content)
	return
}

// 构造实例.
func (o *termFormatter) init() *termFormatter {
	return o
}
