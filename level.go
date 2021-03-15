// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

type Level int

func (c Level) Text() (text string) {
	switch c {
	case LevelError:
		text = "ERROR"
	case LevelWarn:
		text = "Warn"
	case LevelInfo:
		text = "INFO"
	case LevelDebug:
		text = "DEBUG"
	}
	return
}

const (
	LevelOff Level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

var Levels = map[Level]string{
	LevelOff:   "OFF",
	LevelError: "ERROR",
	LevelWarn:  "WARN",
	LevelInfo:  "INFO",
	LevelDebug: "DEBUG",
}
