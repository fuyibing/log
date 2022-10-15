// author: wsfuyibing <websearch@163.com>
// date: 2022-06-02

package log

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
		Client = (&client{}).init()
		Config = (&configuration{}).init()

		Client.Start()
	})
}
