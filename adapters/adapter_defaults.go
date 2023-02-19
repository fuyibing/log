// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package adapters

import (
	"github.com/fuyibing/log/v8/adapters/file"
	"github.com/fuyibing/log/v8/adapters/kafka"
	"github.com/fuyibing/log/v8/adapters/term"
	"github.com/fuyibing/log/v8/formatters"
)

const (
	AdapterError = "error"
	AdapterFile  = "file"
	AdapterTerm  = "term"
	AdapterKafka = "kafka"
)

var (
	adapterContainers = map[string]func() AdapterRegistry{
		AdapterError: func() (ar AdapterRegistry) {
			ar = term.New()
			ar.SetFormatter(formatters.NewTermFormatter())
			return
		},
		AdapterFile: func() (ar AdapterRegistry) {
			ar = file.New()
			ar.SetFormatter(formatters.NewFileFormatter())
			return
		},
		AdapterTerm: func() (ar AdapterRegistry) {
			ar = term.New()
			ar.SetFormatter(formatters.NewTermFormatter())
			return
		},
		AdapterKafka: func() (ar AdapterRegistry) {
			ar = kafka.New()
			ar.SetFormatter(formatters.NewJsonFormatter())
			return
		},
	}
)
