// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/fuyibing/log/v3"
)

func TestClient(t *testing.T) {

	// formatters.Formatter.SetTermFormatter(func(line *base.Line, err error) string {
	//     return "my term:" + line.Content
	// })

	ctx := log.NewContext()

	time.Sleep(time.Second)

	defer func() {
		if r := recover(); r != nil {
			log.Client.Panicfc(ctx, fmt.Sprintf("%v", r))
		}

		time.Sleep(time.Second)
	}()

	log.Client.Infofc(ctx, "Info message")

	ctc := log.ChildContext(ctx, "child context")
	for i := 0; i < 1000; i++ {
		log.Client.Warnfc(ctc, "Warn message: index=%d", i)
	}

	log.Client.Errorfc(ctx, "Error message")
	log.Client.Panicfc(ctx, "not panic")

	time.Sleep(time.Minute)
}
