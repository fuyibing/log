// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

import (
	"fmt"
	"os"
)

var (
	// 1:红, 2:绿, 3:黄, 4: 蓝
	// 5:粉, 6:青, 7:灰
	termColors = map[Level][]int{
		LevelDebug: {30, 47},
		LevelInfo:  {37, 44},
		LevelWarn:  {31, 43},
		LevelError: {33, 41},
	}
)

// Terminal struct.
type adapterTerm struct{}

// New terminal adapter instance.
func NewAdapterTerm() *adapterTerm {
	return new(adapterTerm)
}

// Terminal callback.
func (o *adapterTerm) Callback(line LineInterface) {
	_, _ = fmt.Fprintf(os.Stdout, o.format(line))
}

// Render log content.
func (o *adapterTerm) format(line LineInterface) string {
	// Basic.
	// 1. time.
	// 2. level.
	var s = fmt.Sprintf("[%s][%s][%s][%s:%d]", line.Timeline(), o.level(line), Config.ServerName, Config.ServerAddr, Config.ServerPort)
	// Server info
	// Tracing.
	// 1. version.
	if t := line.Tracing(); t != nil {
		s += fmt.Sprintf("[v=%s] ", t.Version(line.TracingOffset()))
	}
	// Content.
	s += line.String() + "\n"
	return s
}

// Parse log level.
// mixed level name and color.
func (o *adapterTerm) level(line LineInterface) string {
	s := line.GetLevelText()
	if c, ok := termColors[line.GetLevel()]; ok {
		return fmt.Sprintf("%c[%d;%d;%dm%5s%c[0m",
			0x1B, 0,
			c[1], c[0],
			s,
			0x1B,
		)
	} else {
		return fmt.Sprintf("%5s", s)
	}
}
