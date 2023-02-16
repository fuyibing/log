// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package conf

import (
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/adapters/file"
	"github.com/fuyibing/log/v8/adapters/kafka"
	"github.com/fuyibing/log/v8/adapters/term"
)

type (
	Adapter string
	Level   string
)

const (
	Term  Adapter = "term"
	File  Adapter = "file"
	Kafka Adapter = "kafka"
)

func (a Adapter) New() adapters.AdapterInterface {
	switch a {
	case Term:
		return (&term.Executor{}).Init()
	case File:
		return (&file.Executor{}).Init()
	case Kafka:
		return (&kafka.Executor{}).Init()
	}
	return nil
}

const (
	Off   Level = "OFF"
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
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
