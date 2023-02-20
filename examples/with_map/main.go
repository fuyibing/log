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

	log.Map{
		"duration": 3.1415,
		"key":      "value",
	}.Infof("map string")
}
