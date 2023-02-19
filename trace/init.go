// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package trace

import (
	"context"
	"sync"
)

var root context.Context

const (
	StartVersion = "0"
)

func init() {
	new(sync.Once).Do(func() {
		root = context.Background()
	})
}
