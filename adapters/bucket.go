// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package adapters

import (
	"github.com/fuyibing/log/v8/base"
	"sync"
)

type (
	Bucket interface {
		Count() int
		IsEmpty() bool
		Pop() *base.Line
		Popn(limit int) (list []*base.Line, count, remaining int)
		Push(lines ...*base.Line) (total int)
	}

	bucket struct {
		caches []*base.Line
		mu     *sync.Mutex
	}
)

func NewBucket() Bucket {
	return &bucket{
		caches: make([]*base.Line, 0),
		mu:     &sync.Mutex{},
	}
}

func (o *bucket) Count() int {
	o.mu.Lock()
	defer o.mu.Unlock()
	return len(o.caches)
}

func (o *bucket) IsEmpty() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return len(o.caches) == 0
}

func (o *bucket) Pop() *base.Line {
	if list, count, _ := o.Popn(1); count > 0 {
		return list[0]
	}
	return nil
}

func (o *bucket) Popn(limit int) (list []*base.Line, count, remaining int) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Count total items, return if zero
	// returned.
	total := len(o.caches)
	if total == 0 {
		return
	}

	// Execute pop count.
	if total <= limit {
		count = total
		remaining = 0
	} else {
		count = limit
		remaining = total - count
	}

	// Pop slice list.
	list = o.caches[0:count]

	// Reset remaining.
	if remaining > 0 {
		o.caches = o.caches[count:]
	} else {
		o.caches = make([]*base.Line, 0)
	}
	return
}

func (o *bucket) Push(lines ...*base.Line) (total int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.caches = append(o.caches, lines...)
	total = len(o.caches)
	return
}
