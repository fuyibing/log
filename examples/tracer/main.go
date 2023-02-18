// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/conf"
)

func init() {
	log.Config.Set(
		conf.SetAdapter("term"),
		conf.SetLevel(conf.Debug),
		conf.SetTermColor(true),
	)

	// Start log client and accept request.
	log.Client.Start()
}

func main() {
	// Stop log client. Ensure that all data in the
	// memory queue is processed.
	defer log.Client.Stop()

	c1 := log.NewContext()
	log.Debugfc(c1, "example 1 debug")
	log.Infofc(c1, "example 2 info")

	c2 := log.NewChild(c1)
	log.Warnfc(c2, "example 3 warn")
	log.Errorfc(c2, "example 4 error")
	log.Panicfc(c2, "example 5 panic")
}
