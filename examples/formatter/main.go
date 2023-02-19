// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/conf"
)

func init() {
	log.Config.Set(
		conf.SetAdapter(adapters.AdapterTerm),
		// conf.SetAdapter(adapters.AdapterFile),
		conf.SetLevel(conf.Debug),
		conf.SetTermColor(true),

		conf.SetServiceAddr("172.16.0.100"),
		conf.SetServiceEnvironment("production"),
		conf.SetServiceName("mbs"),
		conf.SetServiceVersion("3.1.2"),
		conf.SetServicePort(8101),
	)

	// Start log client and accept request.
	log.Client.Reset()

	// Override default formatter.
	// log.Client.GetAdapterRegistry().SetFormatter(formatters.NewJsonFormatter())
}

func main() {
	// Close log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Close()

	log.Debug("example 1 debug")
	log.Info("example 2 info")
	log.Warn("example 3 warn")
	log.Error("example 4 error")
	log.Panic("example 5 panic")
}
