// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package formatters

import "sync"

func init() {
	new(sync.Once).Do(func() {
		Formatter = (&formatter{}).init()

		dataPool = &sync.Pool{
			New: func() interface{} {
				return (&data{}).init()
			},
		}
	})
}
