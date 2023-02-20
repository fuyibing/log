// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/formatters"
)

func init() {
	// log.Client.Reset()
	log.Client.GetAdapterRegistry().SetFormatter(&formatters.FileFormatter{})
}

func main() {
	// Close log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Close()

	log.Debug("debug")
	log.Infof("text format, id=%d, type=%s", 100, "demo")
	log.Map{"id": 100, "type": "demo"}.Infof("with map config")

	c1 := log.NewContext()
	log.Debugfc(c1, "debug")
	log.Infofc(c1, "text format, id=%d, type=%s", 100, "demo")
	log.Map{"id": 100, "type": "demo"}.Infofc(c1, "with map config")

	c2 := log.NewChild(c1)
	log.Infofc(c2, "child of map 1")
	log.Infofc(c2, "child of map 2")
}
