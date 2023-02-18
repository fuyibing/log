// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"github.com/fuyibing/log/v8/base"
	"sync"
)

type (
	// Bucket
	// memory queue manager.
	Bucket interface {
		// Count
		// return total lines in memory queue.
		Count() int

		// Pop
		// pop one line from memory queue.
		Pop() (line *base.Line, exists bool)

		// Popn
		// pop specified count lines from memory queue.
		Popn(limit int) (list []*base.Line, total, count, remaining int)

		// Push
		// lines into memory queue.
		Push(lines ...*base.Line) (total, count int)
	}

	bucket struct {
		mu     *sync.Mutex
		queues []*base.Line
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *bucket) Count() int {
	o.mu.Lock()
	defer o.mu.Unlock()

	return len(o.queues)
}

func (o *bucket) Pop() (line *base.Line, exists bool) {
	if list, _, count, _ := o.Popn(1); count > 0 {
		line = list[0]
		exists = true
	}
	return
}

func (o *bucket) Popn(limit int) (list []*base.Line, total, count, remaining int) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Do nothing
	// if queue is empty.
	if total = len(o.queues); total == 0 {
		return
	}

	// Pop all lines.
	if total <= limit {
		count = total
		remaining = 0
		list = o.queues[0:]
		o.queues = make([]*base.Line, 0)
		return
	}

	// Pop slice.
	count = limit
	remaining = total - count
	list = o.queues[0:count]
	o.queues = o.queues[count:]
	return
}

func (o *bucket) Push(lines ...*base.Line) (total, count int) {
	o.mu.Lock()
	defer o.mu.Unlock()

	total = len(o.queues)

	if count = len(lines); count > 0 {
		o.queues = append(o.queues, lines...)
		total += count
	}
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *bucket) init() *bucket {
	o.mu = &sync.Mutex{}
	o.queues = make([]*base.Line, 0)
	return o
}
