// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package tests

import (
	"sync"
	"testing"
	"time"

	"github.com/fuyibing/log/v3"
)

func TestClient(t *testing.T) {

	ct1 := log.NewContext()
	log.Client.Debugfc(ct1, "debug 1")
	log.Client.Infofc(ct1, "info 2")

	time.Sleep(time.Second)
	ct2 := log.ChildContext(ct1, "child span")
	log.Client.Infofc(ct2, "child info 1")
	log.Client.Warnfc(ct2, "child warn 2")

	log.Client.Warnfc(ct1, "warn 3")
	log.Client.Errorfc(ct1, "error 4")

	time.Sleep(time.Second)
}

func TestTryLock(t *testing.T) {

	mu := sync.Mutex{}

	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			mu.Lock()
			defer func() {
				mu.Unlock()
				wg.Done()
			}()

			println("begin:", index)
			time.Sleep(time.Second)
			println("end:", index)
		}(i)
	}

	wg.Wait()

}
