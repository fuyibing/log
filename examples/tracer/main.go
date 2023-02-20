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
	log.Client.Reset()
}

func main() {
	// Close log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Close()

	c1 := log.NewContext()
	log.Debugfc(c1, "example 1 debug")
	log.Infofc(c1, "example 2 info")

	c2 := log.NewChild(c1)
	log.Map{"key": "value", "task-id": 100}.Infofc(c2, "with map")
	log.Warnfc(c2, "example 3 warn")
	log.Errorfc(c2, "example 4 error")
	log.Panicfc(c2, "example 5 panic")
}
