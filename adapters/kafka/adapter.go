// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package kafka

import (
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formatters"
	"sync"
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

func (o *adapter) doLogs(lines ...*base.Line) {}

func (o *adapter) init() *adapter {
	o.formatter = formatters.NewKafkaFormatter().Format
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
