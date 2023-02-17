// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package base

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
		Pool = (&pool{}).init()
	})
}
