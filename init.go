// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package log

import (
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/managers"
	"sync"
)

var (
	Config *conf.Configuration
	Client *managers.Client
)

func init() {
	new(sync.Once).Do(func() {
		Config = conf.Config
		Client = (&managers.Client{}).Init()
		Client.Subscribe()
	})
}
