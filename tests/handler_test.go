// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package tests

import (
	"testing"

	"github.com/fuyibing/log"
	"github.com/fuyibing/log/interfaces"
)

func TestHandler(t *testing.T) {

	log.Config.SetHandler(func(line interfaces.LineInterface) {
		println("handler: ", line.SpanVersion(), line.Content())
	})

	ctx := log.NewContext()

	log.Client.Debugfc(ctx, "debug fc")
	log.Client.Infofc(ctx, "info fc")
	log.Client.Warnfc(ctx, "warn fc")
	log.Client.Errorfc(ctx, "error fc")
}
