// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package term

import (
	"fmt"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
	"os"
	"strings"
	"sync"
)

var (
	Colors = map[conf.Level][]int{
		conf.Debug: {30, 0}, // Text: black, Background: white
		conf.Info:  {32, 0}, // Text: green, Background: white
		conf.Warn:  {33, 0}, // Text: yellow, Background: white
		conf.Error: {31, 0}, // Text: red, Background: white
		conf.Fatal: {31, 0}, // Text: red, Background: white
	}
)

type (
	adapter struct {
		child     adapters.AdapterManager
		formatter formatters.Formatter
		ignorer   adapters.AdapterIgnore
	}
)

// /////////////////////////////////////////////////////////////
// Exported interface methods
// /////////////////////////////////////////////////////////////

func (o *adapter) Log(lines ...*base.Line)             { o.doLogs(lines...) }
func (o *adapter) SetChild(v adapters.AdapterManager)  { o.child = v }
func (o *adapter) SetFormatter(v formatters.Formatter) { o.formatter = v }
func (o *adapter) SetIgnore(v adapters.AdapterIgnore)  { o.ignorer = v }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *adapter) doLogs(lines ...*base.Line) {
	list := make([]string, 0)

	// Generate contents
	for _, x := range lines {
		list = append(list, func(line *base.Line) string {
			if conf.Config.Term.Color {
				if c, ok := Colors[line.Level]; ok {
					return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, c[1], c[0], o.formatter(line), 0x1B)
				}
			}
			return o.formatter(line)
		}(x))

		x.Release()
	}

	// Send to stdout.
	_, _ = fmt.Fprintf(os.Stdout,
		fmt.Sprintf("%s\n",
			strings.Join(list, "\n"),
		),
	)
}

func (o *adapter) init() *adapter {
	o.formatter = formatters.NewTextFormatter().Format
	return o
}

// /////////////////////////////////////////////////////////////
// Package init
// /////////////////////////////////////////////////////////////

var (
	Adapter *adapter
)

func init() {
	new(sync.Once).Do(func() {
		Adapter = (&adapter{}).init()
	})
}
