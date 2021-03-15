// author: wsfuyibing <websearch@163.com>
// date: 2021-03-15

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2"
)

func TestTracing(t *testing.T) {
	x := log.NewTracing().UseDefault()
	for n := 0; n < 3; n++ {
		id, nid := x.Increment()
		t.Logf("current=%d and next=%d and ver=%s and link=%s.", id, nid, x.Version(id), x.LinkVersion())
	}
}
