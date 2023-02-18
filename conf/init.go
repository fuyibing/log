// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

import (
	"sync"
)

var Config Configuration

func init() {
	new(sync.Once).Do(func() {
		Config = func() Configuration {
			c := (&configuration{}).init()
			c.initYaml()
			c.initDefaults()
			return c
		}()
	})
}
