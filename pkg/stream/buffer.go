package stream

import (
	"sync/atomic"
)

type RingBuffer struct {
	buffer []interface{}
	size   uint64
	head   uint64
	tail   uint64
}

func NewRingBuffer(size uint64) *RingBuffer {
	return &RingBuffer{
		buffer: make([]interface{}, size),
		size:   size,
	}
}

func (r *RingBuffer) Push(val interface{}) bool {
	head := atomic.LoadUint64(&r.head)
	tail := atomic.LoadUint64(&r.tail)
	if head-tail >= r.size {
		return false
	}
	r.buffer[head%r.size] = val
	atomic.AddUint64(&r.head, 1)
	return true
}

func (r *RingBuffer) Pop() (interface{}, bool) {
	tail := atomic.LoadUint64(&r.tail)
	head := atomic.LoadUint64(&r.head)
	if tail >= head {
		return nil, false
	}
	val := r.buffer[tail%r.size]
	atomic.AddUint64(&r.tail, 1)
	return val, true
}
