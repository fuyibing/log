// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

import (
	"sync"
)

var (
	Config *Configuration
)

func init() {
	new(sync.Once).Do(func() {
		Config = (&Configuration{}).init()
	})
}
