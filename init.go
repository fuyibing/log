// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/core"
	"sync"
)

var (
	Config conf.Configuration
	Client core.Client
)

func init() {
	new(sync.Once).Do(func() {
		Config = conf.Config
		Client = core.NewClient()

		// Auto start if enabled.
		if Config.GetAutoStart() {
			Client.Reset()
		}
	})
}
