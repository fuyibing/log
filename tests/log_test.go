// author: wsfuyibing <websearch@163.com>
// date: 2021-02-11

package tests

import (
	"testing"
	"time"

	"github.com/fuyibing/log"
)

func TestLogTerm(t *testing.T) {
	ctx := log.NewContext()

	l := log.New()
	l.SetLevel(log.LevelDebug)
	l.SetAdapter(log.AdapterTerm)

	l.Debugfc(ctx, "Debug format context")
	l.Infofc(ctx, "Info format context")
	l.Warnfc(ctx, "Warn format context")
	l.Errorfc(ctx, "Error format context")

	time.Sleep(time.Second)
}

func TestLogRedis(t *testing.T) {
	ctx := log.NewContext()

	l := log.New()
	l.SetLevel(log.LevelDebug)
	l.SetAdapter(log.AdapterRedis)

	l.Debugfc(ctx, "Debug format context")
	l.Infofc(ctx, "Info format context")
	l.Warnfc(ctx, "Warn format context")
	l.Errorfc(ctx, "Error format context")

	time.Sleep(time.Second)
}
