// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package formatters

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

type (
	IgnoreFormatter struct {
	}
)

func NewIgnoreFormatter() *IgnoreFormatter {
	return &IgnoreFormatter{}
}

func (o *IgnoreFormatter) Format(line *base.Line) string {
	return fmt.Sprintf("<?> %s[%s][%s]",
		conf.Config.Prefix, line.Time.Format(conf.Config.TimeFormat), line.Level.String(),
	) + " " + line.Text
}
