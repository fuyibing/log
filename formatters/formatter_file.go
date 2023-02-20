// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"encoding/json"
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

type (
	FileFormatter struct{}
)

func NewFileFormatter() *FileFormatter {
	return &FileFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *FileFormatter) Body(line *base.Line) []byte {
	return []byte(o.String(line))
}

// String
// generate log line as text.
//
// About Keyword.
//   D  : Duration
//   P  : Process id of os
//   SN : Registered service name
//   T  : Trace id
//   TS : Span id of trace
//   TP : Parent Span id of trace
//   TV : Span version
//   R  : HTTP Request URL
//   RM : HTTP Request Method
func (o *FileFormatter) String(line *base.Line) (str string) {
	// Prefix.
	str = conf.Config.GetPrefix()

	// Time & Level.
	str += fmt.Sprintf("[%s][%s]",
		line.Time.Format(conf.Config.GetTimeFormat()),
		line.Level,
	)

	// Service: host + port.
	if s := conf.Config.GetServiceAddr(); s != "" {
		str += fmt.Sprintf("[%s:%d]", s, conf.Config.GetServicePort())
	}

	// PID.
	str += fmt.Sprintf("{P=%d}", conf.Config.GetPid())

	// Duration.
	if line.Duration > 0 {
		str += fmt.Sprintf("{D=%v}", line.Duration)
	}

	// Service: name.
	if s := conf.Config.GetServiceName(); s != "" {
		str += fmt.Sprintf("{SN=%s}", s)
	}

	// Open Tracing.
	if t := line.Tracing(); t != nil {
		str += fmt.Sprintf("{T=%s}{TS=%s}{TP=%s}{TV=%v}",
			t.TraceId, t.SpanId, t.ParentSpanId,
			t.GenVersion(line.TracingOffset()),
		)

		// Append http request location.
		if t.Http {
			str += fmt.Sprintf("{R=%s}{RM=%s}",
				t.HttpRequestUrl,
				t.HttpRequestMethod,
			)
		}
	}

	// Property.
	if line.Property != nil {
		if buf, err := json.Marshal(line.Property); err == nil {
			str += fmt.Sprintf(" %s", buf)
		}
	}

	// User info.
	str += " " + line.Text
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *FileFormatter) init() *FileFormatter {
	return o
}
