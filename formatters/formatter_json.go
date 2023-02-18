// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"encoding/json"
	"github.com/fuyibing/log/v8/base"
)

type (
	JsonFormatter struct{}

	// JsonLine
	// for kafka.
	JsonLine struct {
		Datetime int64 `json:"datetime"`

		Content string `json:"content"`
	}
)

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) Body(line *base.Line) []byte {
	body, _ := json.Marshal((&JsonLine{}).init(line))
	return body
}

func (o *JsonFormatter) String(line *base.Line) string {
	return string(o.Body(line))
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) init() *JsonFormatter { return o }

func (o *JsonLine) init(line *base.Line) *JsonLine {
	o.Datetime = line.Time.UnixMilli()

	o.Content = line.Text
	return o
}
