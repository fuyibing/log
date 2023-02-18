// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"testing"
)

func BenchmarkBucket_Count(b *testing.B) {
	x := bucketCreate()
	for i := 0; i < b.N; i++ {
		x.Count()
	}
}

func BenchmarkBucket_Pop(b *testing.B) {
	x := bucketCreate()
	for i := 0; i < b.N; i++ {
		x.Pop()
	}
}

func BenchmarkBucket_Popn(b *testing.B) {
	x := bucketCreate()
	for i := 0; i < b.N; i++ {
		x.Popn(1)
	}
}

func BenchmarkBucket_Push(b *testing.B) {
	x := bucketCreate()
	for i := 0; i < b.N; i++ {
		x.Push(&base.Line{Text: fmt.Sprintf("example %d", i)})
	}
}

func bucketCreate() Bucket {
	x := (&bucket{}).init()
	for i := 0; i < 1000; i++ {
		x.Push(&base.Line{Text: fmt.Sprintf("example: %d", i)})
	}
	return x
}
