// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import "sync"

func init() {
	new(sync.Once).Do(func() {
		linePool = &sync.Pool{
			New: func() interface{} {
				return (&Line{}).init()
			},
		}
	})
}
