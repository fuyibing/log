// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package base

type (
	Line struct {
	}
)

// Release
// put this to pool.
func (o *Line) Release() { Pool.ReleaseLine(o) }

// /////////////////////////////////////////////////////////////
// Constructor, set, unset methods
// /////////////////////////////////////////////////////////////

func (o *Line) after() *Line { return o }

func (o *Line) before() *Line { return o }

func (o *Line) init() *Line { return o }
