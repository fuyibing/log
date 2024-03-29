// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: wsfuyibing <websearch@163.com>
// date: 2023-02-28

package common

import (
	"fmt"
	"sync"
)

var (
	ErrBucketIsFull = fmt.Errorf("bucket is fully")
)

type (
	// Bucket component used for memory queues. In this package, his role used
	// to store loggers.Log and tracers.Span components. Async goroutines pop
	// them then publish to specified adapters (eg. Kafka, Jaeger and so on).
	Bucket interface {
		// Add
		// items (loggers.Log / tracers.Span) into memory queue.
		Add(item interface{}) (total int, err error)

		// Count
		// return total items in bucket.
		Count() int

		// IsEmpty
		// return true if no item in bucket.
		IsEmpty() bool

		// Pop
		// pop one item from bucket.
		Pop() (item interface{}, exists bool)

		// Popn
		// pop specified count items from bucket.
		Popn(limit int) (items []interface{}, total, count int)

		// SetCapacity
		// config bucket capacity.
		SetCapacity(n int) Bucket
	}

	bucket struct {
		sync.Mutex

		caches   []interface{}
		capacity int
	}
)

func NewBucket(capacity int) Bucket { return (&bucket{capacity: capacity}).init() }

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *bucket) Add(item interface{}) (total int, err error) {
	o.Lock()
	defer o.Unlock()

	if item != nil {
		if total = len(o.caches) + 1; o.capacity > 0 && total > o.capacity {
			err = ErrBucketIsFull
			return
		}
	}

	o.caches = append(o.caches, item)
	return
}

func (o *bucket) Count() int {
	o.Lock()
	defer o.Unlock()

	return len(o.caches)
}

func (o *bucket) IsEmpty() bool {
	return 0 == o.Count()
}

func (o *bucket) Pop() (item interface{}, exists bool) {
	if items, _, count := o.Popn(1); count == 1 {
		item = items[0]
		exists = true
	}
	return
}

func (o *bucket) Popn(limit int) (items []interface{}, total, count int) {
	o.Lock()
	defer o.Unlock()

	if total = len(o.caches); total == 0 {
		return
	}

	if limit >= total {
		count = total
		items = o.caches[:]
		o.reset()
		return
	}

	count = limit
	items = o.caches[0:count]
	o.caches = o.caches[count:]
	return
}

func (o *bucket) SetCapacity(n int) Bucket {
	o.Lock()
	defer o.Unlock()

	o.capacity = n
	return o
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *bucket) init() *bucket {
	o.reset()
	return o
}

func (o *bucket) reset() *bucket {
	o.caches = make([]interface{}, 0)
	return o
}
