// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

type (
	Level string
)

const (
	Fatal Level = "FATAL"
	Error Level = "ERROR"
	Warn  Level = "WARN"
	Info  Level = "INFO"
	Debug Level = "DEBUG"
)

func (l Level) Int() int {
	switch l {
	case Fatal:
		return 1
	case Error:
		return 2
	case Warn:
		return 3
	case Info:
		return 4
	case Debug:
		return 5
	}
	return 0
}

func (l Level) String() string {
	return string(l)
}
