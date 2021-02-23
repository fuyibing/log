// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package tests

import (
	"testing"
	"time"

	"github.com/fuyibing/log/v2"
)

func TestFunc(t *testing.T) {
	// ctx := log.NewContext()
	ctx := log.NewTracing().UseDefault()

	// println("traceid:", ctx.Value(interfaces.OpenTracingKey).(interfaces.TraceInterface).GetTraceId())

	log.Client.Debugfc(ctx, "debug fc")
	log.Client.Infofc(ctx, "info fc")
	log.Client.Warnfc(ctx, "warn fc")
	log.Client.Errorfc(ctx, "error fc")

	time.Sleep(time.Second * 60)
}
