// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

type (
	// TermFormatter
	// use to print on console/terminal.
	TermFormatter struct{}
)

func NewTermFormatter() *TermFormatter {
	return &TermFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *TermFormatter) Body(line *base.Line) []byte { return []byte(o.String(line)) }

func (o *TermFormatter) String(line *base.Line) (str string) {
	// Prefix.
	// Use fixed time-format and length
	str = fmt.Sprintf("%s[%-15s][%5s]", conf.Config.GetPrefix(),
		line.Time.Format("15:04:05.999999"),
		line.Level,
	)

	if line.Duration > 0 {
		str += fmt.Sprintf("{D=%v}", line.Duration)
	}

	// Open Tracing.
	if t := line.Tracing(); t != nil {
		str += fmt.Sprintf("{%s=%v}", t.SpanId, t.GenVersion(line.TracingOffset()))
		if t.Http {
			str += fmt.Sprintf("{%s=%s}", t.HttpRequestMethod, t.HttpRequestUrl)
		}
	}

	// User info.
	str += " " + line.Text
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *TermFormatter) init() *TermFormatter {
	return o
}
