// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package tests

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/conf"
)

func init() {
	// log.Config.Set(conf.WithAdapter(conf.Term, conf.File, conf.Kafka))
	// log.Config.Set(conf.WithAdapter(conf.Term), conf.WithTimeFormat("15:04:05.999999"))
	log.Config.Set(conf.WithAdapter(conf.File), conf.WithTimeFormat("15:04:05.999999"))
	// log.Config.Set(conf.WithAdapter(), conf.WithTimeFormat("15:04:05.999999"))
	log.Config.Term.Color = true

	log.Client.Start(nil)
}
