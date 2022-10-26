// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import "sync"

func init() {
	new(sync.Once).Do(func() {
		Client = &ClientManager{}
		Client.init()

		Config = &Configuration{}
		Config.init()
	})
}
