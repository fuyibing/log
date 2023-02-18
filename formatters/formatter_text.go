// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

type (
	TextFormatter struct{}
)

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *TextFormatter) Body(line *base.Line) []byte {
	return []byte(o.String(line))
}

func (o *TextFormatter) String(line *base.Line) (str string) {
	// Prefix.
	str = conf.Config.GetPrefix()

	// Time & Level & PID.
	str += fmt.Sprintf("[%s][%s][PI=%d]",
		line.Time.Format(conf.Config.GetTimeFormat()),
		line.Level,
		conf.Config.GetPid(),
	)

	// Append service name.
	if s := conf.Config.GetServiceName(); s != "" {
		str += fmt.Sprintf("[SN=%s]", s)
	}

	// Open Tracing.
	if t := line.Tracing(); t != nil {
		str += fmt.Sprintf("[T=%s][TS=%s][TP=%s][TV=%v]",
			t.TraceId,
			t.SpanId,
			t.ParentSpanId,
			t.GenVersion(line.TracingOffset()),
		)
	}

	// User info.
	str += " " + line.Text
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *TextFormatter) init() *TextFormatter {
	return o
}
