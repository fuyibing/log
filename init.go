// author: wsfuyibing <websearch@163.com>
// date: 2022-06-02

package log

import (
	"sync"
)

func init() {
	new(sync.Once).Do(func() {
		// 1. 构建实例.
		Client = (&client{}).init()
		Config = (&Configuration{}).init()

		// 2. 启动客户端.
		Client.Start()
	})
}
