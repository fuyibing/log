// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package formatters

import (
	"fmt"

	"github.com/fuyibing/log/v3/base"
)

// File格式化.
// 格式化后的数据用于写入文件存储中.
type fileFormatter struct{}

// Format
// 格式化过程.
func (o *fileFormatter) Format(line *base.Line, err error) (text string) {
	// 1. 基础信息.
	text = fmt.Sprintf("[%s][%s:%d][%s][%s][PID=%d]",
		line.Time.Format(base.LogTimeFormat),
		base.LogHost, base.LogPort,
		base.LogName, line.Level.Name(),
		base.LogPid,
	)

	// 2. 链路信息.
	if line.Trace {
		text += fmt.Sprintf("[TRACE-ID=%s][PARENT-SPAN-ID=%s][SPAN-ID=%s][SPAN-VERSION=%s.%d]",
			line.TraceId, line.ParentSpanId,
			line.SpanId, line.SpanPrefix, line.SpanOffset,
		)
		if line.Duration > 0 {
			text += fmt.Sprintf("[DURATION=%.06f]", line.Duration)
		}
		if line.RequestMethod != "" {
			text += fmt.Sprintf("[REQUEST-METHOD=%s]", line.RequestMethod)
		}
		if line.RequestUrl != "" {
			text += fmt.Sprintf("[REQUEST-URL=%s]", line.RequestUrl)
		}
	}

	// 3. 基础内容.
	text += fmt.Sprintf(" %s", line.Content)
	if err != nil {
		text += fmt.Sprintf(" << interrpt: %s", err.Error())
	}
	return
}

// 构造实例.
func (o *fileFormatter) init() *fileFormatter {
	return o
}
