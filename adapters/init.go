// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package adapters

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
		linePool = &sync.Pool{
			New: func() interface{} {
				return (&Line{}).init()
			},
		}
	})
}
