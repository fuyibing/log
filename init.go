// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package log

import (
	"sync"

	"github.com/fuyibing/log/interfaces"
)

var (
	Config interfaces.ConfigInterface
	Client interfaces.ClientInterface
)

func init() {
	new(sync.Once).Do(func() {
		Config = newConfiguration()
		Client = newClient()
	})
}
