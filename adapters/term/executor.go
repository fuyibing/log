// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package term

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
	"os"
	"strings"
)

var (
	colors = map[conf.Level][]int{
		conf.Debug: {30, 0}, // Text: black, Background: white
		conf.Info:  {32, 0}, // Text: green, Background: white
		conf.Warn:  {33, 0}, // Text: yellow, Background: white
		conf.Error: {31, 0}, // Text: red, Background: white
		conf.Fatal: {31, 0}, // Text: red, Background: white
	}
)

type (
	// Executor
	// write log content to terminal, print on console.
	Executor struct {
		formatter formatters.Formatter
	}
)

func New() *Executor {
	return (&Executor{}).init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Executor) Logs(lines ...*base.Line) (err error) {
	list := make([]string, 0)

	// Iterate lines.
	for _, line := range lines {
		list = append(list, o.format(line))
	}

	// Send to standard output.
	_, err = fmt.Fprintf(os.Stdout,
		fmt.Sprintf("%s\n", strings.Join(list, "\n")),
	)

	return
}

func (o *Executor) SetFormatter(formatter formatters.Formatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Executor) format(line *base.Line) string {
	if conf.Config.GetTerm().GetColor() {
		if c, ok := colors[line.Level]; ok {
			return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
				0x1B, 0, c[1], c[0], o.formatter.String(line), 0x1B,
			)
		}
	}

	return o.formatter.String(line)
}

func (o *Executor) init() *Executor {
	return o
}
