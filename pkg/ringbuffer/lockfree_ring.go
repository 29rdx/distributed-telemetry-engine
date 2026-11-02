package ringbuffer

import (
	"sync/atomic"
)

type RingBuffer struct {
	data     []interface{}
	capacity uint64
	mask     uint64
	head     uint64
	tail     uint64
}

func NewRingBuffer(powerOfTwo uint64) *RingBuffer {
	capacity := uint64(1) << powerOfTwo
	return &RingBuffer{
		data:     make([]interface{}, capacity),
		capacity: capacity,
		mask:     capacity - 1,
	}
}

func (r *RingBuffer) Push(val interface{}) bool {
	for {
		head := atomic.LoadUint64(&r.head)
		tail := atomic.LoadUint64(&r.tail)
		if tail-head >= r.capacity {
			return false // Ring buffer full
		}
		if atomic.CompareAndSwapUint64(&r.tail, tail, tail+1) {
			r.data[tail&r.mask] = val
			return true
		}
	}
}

func (r *RingBuffer) Pop() (interface{}, bool) {
	for {
		head := atomic.LoadUint64(&r.head)
		tail := atomic.LoadUint64(&r.tail)
		if head == tail {
			return nil, false // Ring buffer empty
		}
		if atomic.CompareAndSwapUint64(&r.head, head, head+1) {
			val := r.data[head&r.mask]
			return val, true
		}
	}
}
