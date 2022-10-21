// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package tests

import (
	"github.com/fuyibing/log/v3"
	"github.com/fuyibing/log/v3/trace"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	t.Log("------------ logger begin ---------")

	defer func() {
		time.Sleep(time.Minute * 1)
		log.Client.Stop()
		t.Log("------------ logger quited ---------")
	}()

	time.Sleep(time.Second)

	ctx := trace.New()
	num := time.Now().Nanosecond()
	for i := 0; i < 10; i++ {
		log.Client.Infofc(ctx, "info %d.%d", num, i)
	}
}
