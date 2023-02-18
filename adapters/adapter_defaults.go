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
	adapterFile  = "file"
	adapterTerm  = "term"
	adapterKafka = "kafka"
)

var (
	adapterDefaults = map[string]func() AdapterRegistry{
		adapterFile: func() (ar AdapterRegistry) {
			ar = file.New()
			ar.SetFormatter(formatters.NewTextFormatter())
			return
		},
		adapterTerm: func() (ar AdapterRegistry) {
			ar = term.New()
			ar.SetFormatter(formatters.NewTextFormatter())
			return
		},
		adapterKafka: func() (ar AdapterRegistry) {
			ar = kafka.New()
			ar.SetFormatter(formatters.NewJsonFormatter())
			return
		},
	}
)
