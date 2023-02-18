// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"encoding/json"
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/conf"
	"time"
)

func init() {
	log.Config.Set(
		conf.SetAsyncDisabled(true),
		conf.SetTimeFormat("05.999"),
		conf.SetServiceHost("127.0.0.1"),
		conf.SetServicePort(8080),
	)
	log.Client.Reset()
}

func main() {
	// Close log client. Ensure that all data in the
	// memory queue are processed.
	defer func() {
		time.Sleep(time.Second * 3)
		log.Client.Close()
	}()

	buf, _ := json.Marshal(log.Config)
	log.Infof("configuration: %s", buf)
}
