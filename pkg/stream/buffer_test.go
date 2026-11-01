package stream

import "testing"

func TestRingBufferPushPop(t *testing.T) {
	rb := NewRingBuffer(16)
	if !rb.Push("metric_payload_1") {
		t.Fatal("expected push to succeed")
	}
	val, ok := rb.Pop()
	if !ok || val != "metric_payload_1" {
		t.Fatalf("unexpected pop value: %v", val)
	}
}
