package ringbuffer

import (
	"testing"
)

func TestRingBufferBasic(t *testing.T) {
	ring := NewRingBuffer(4) // 16 capacity
	if !ring.Push(100) {
		t.Fatal("expected push to succeed")
	}
	val, ok := ring.Pop()
	if !ok || val.(int) != 100 {
		t.Fatalf("expected 100, got %v", val)
	}
}
