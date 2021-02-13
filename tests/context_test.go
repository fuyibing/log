// author: wsfuyibing <websearch@163.com>
// date: 2021-02-11

package tests

import (
	"testing"

	"github.com/fuyibing/log"
)

func TestContextNew(t *testing.T) {
	c := log.NewContext()
	x := c.Value(log.OpenTracingContext).(log.TracingInterface)
	t.Logf("Context: {%v}.", c.Err())
	t.Logf("Tracing: trace id: %s.", x.TraceId())
	t.Logf("         span  id: %s.", x.SpanId())
	t.Logf("         span ver: %s.", x.SpanVersion())
}
