// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
	})
}
