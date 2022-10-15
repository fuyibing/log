// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package redis

import "sync"

func init() {
	new(sync.Once).Do(func() {
		Config = (&Configuration{}).init()
	})
}
