// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

// Custom formatter.
type formatter struct{}

func (*formatter) Body(line *base.Line) []byte { return nil }

func (*formatter) String(line *base.Line) string {
	return "my formatter: " + line.Text
}

func init() {
	log.Config.Set(
		conf.SetAdapter(adapters.AdapterTerm),
		conf.SetLevel(conf.Debug),
		conf.SetTermColor(true),
	)

	// Start log client and accept request.
	log.Client.Reset()

	// Override default formatter.
	log.Client.GetAdapterRegistry().
		SetFormatter(&formatter{})
}

func main() {
	// Stop log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Stop()

	log.Debug("example 1 debug")
	log.Info("example 2 info")
	log.Warn("example 3 warn")
	log.Error("example 4 error")
	log.Panic("example 5 panic")
}
