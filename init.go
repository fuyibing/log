// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import "sync"

func init() {
	new(sync.Once).Do(func() {
		Config = (&configuration{}).init()

		Client = (&client{}).init()
		Client.Start()
	})
}
