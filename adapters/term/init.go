// author: wsfuyibing <websearch@163.com>
// date: 2022-10-14

package term

import "sync"

func init() {
	new(sync.Once).Do(func() {
	})
}
