// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package base

import (
	"sync"
)

var (
	Parser ParserManager
	Pool   PoolManager
)

func init() {
	new(sync.Once).Do(func() {
		Parser = (&parser{}).init()
		Pool = (&pool{}).init()
	})
}
