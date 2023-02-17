// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package managers

import (
	"github.com/fuyibing/log/v8/conf"
	"sync"
)

var Client *ClientManager

func init() {
	new(sync.Once).Do(func() {
		Client = (&ClientManager{}).init()

		if conf.Config.AutoStart {
			Client.Start(nil)
		}
	})
}
