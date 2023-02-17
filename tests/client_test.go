// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package tests

import (
	"github.com/fuyibing/log/v8"
	"sync"
	"testing"
)

func TestClient_Default(t *testing.T) {
	w := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			clientDefault(i)
		}(i)
	}
	w.Wait()

	log.Client.Stop()
}

func clientDefault(i int) {
	log.Debugf("Debug %d.1", i)
	log.Infof("Info %d.2", i)
	log.Warnf("Warn %d.3", i)
	log.Errorf("Error %d.4", i)
	log.Panicf("Panic %d.5", i)
}
