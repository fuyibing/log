// author: wsfuyibing <websearch@163.com>
// date: 2021-02-11

package tests

import (
	"testing"

	"github.com/fuyibing/log"
)

func TestTracing(t *testing.T) {
	x := log.NewTracing().UseDefault()
	t.Logf("Tracing: trace id: %s.", x.TraceId())
	t.Logf("         span  id: %s.", x.SpanId())
	t.Logf("         span ver: %s.", x.SpanVersion())
}
