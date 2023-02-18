// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package adapters

import (
	"sync"
)

var (
	Adapter AdapterManager
)

func init() {
	new(sync.Once).Do(func() {
		Adapter = (&adapter{}).init()
	})
}
