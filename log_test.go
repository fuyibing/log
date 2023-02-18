// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"testing"
)

type testFormatter struct {
}

func (*testFormatter) Body(line *base.Line) []byte { return nil }

func (*testFormatter) String(line *base.Line) string {
	return "my formatter: " + line.Text
}

func TestClient(t *testing.T) {
	Config.Set(
		conf.SetLevel(conf.Debug),
		conf.SetPrefix("# "),
		conf.SetServiceName("v8log"),
		conf.SetTermColor(true),
	)

	Client.Start()
	defer func() {
		Client.Stop()
	}()

	Client.GetAdapterRegistry().SetFormatter(&testFormatter{})

	for i := 0; i < 1; i++ {
		func(index int) {
			c1 := NewContextInfo("top context")
			Debugfc(c1, "debug %d.1", index)
			Infofc(c1, "info %d.2", index)

			c2 := NewChildInfo(c1, "top child")
			Warnfc(c2, "warn %d.3", index)
			Errorfc(c2, "error %d.4", index)
			Fatalfc(c2, "fatal %d.5", index)
		}(i)
	}
}
