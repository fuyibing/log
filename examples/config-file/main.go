// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"encoding/json"
	"github.com/fuyibing/log/v8"
)

func init() {
	log.Client.Reset()
}

func main() {
	// Stop log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Stop()

	buf, _ := json.Marshal(log.Config)
	log.Infof("configuration: %s", buf)
}
