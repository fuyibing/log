// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import "strings"

type (
	Level     int
	LevelName string
)

const (
	Off Level = iota
	Error
	Warn
	Info
	Debug
)

// Name
// 级别编号转名称.
func (l Level) Name() (s string) {
	switch l {
	case Debug:
		s = "DEBUG"
	case Info:
		s = "INFO"
	case Warn:
		s = "WARN"
	case Error:
		s = "ERROR"
	default:
		s = "OFF"
	}
	return
}

// Level
// 级别名称转编号.
func (s LevelName) Level() (l Level) {
	switch strings.ToUpper(string(s)) {
	case "DEBUG":
		l = Debug
	case "INFO":
		l = Info
	case "WARN":
		l = Warn
	case "ERROR":
		l = Error
	default:
		l = Off
	}
	return
}
