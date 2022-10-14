// author: wsfuyibing <websearch@163.com>
// date: 2022-10-14

package redis

import "sync"

func init() {
	new(sync.Once).Do(func() {

		dataPool = &sync.Pool{
			New: func() interface{} {
				return (&Data{}).init()
			},
		}

		Config = (&Configuration{}).init()
	})
}
