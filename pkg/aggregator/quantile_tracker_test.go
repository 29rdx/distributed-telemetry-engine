package aggregator

import (
	"testing"
)

func TestQuantileTracker(t *testing.T) {
	tracker := NewQuantileTracker(100)
	for i := 1; i <= 100; i++ {
		tracker.Record(float64(i))
	}
	p50 := tracker.GetPercentile(0.5)
	p99 := tracker.GetPercentile(0.99)
	if p50 < 49 || p50 > 51 {
		t.Fatalf("expected P50 around 50, got %f", p50)
	}
	if p99 < 98 {
		t.Fatalf("expected P99 around 99, got %f", p99)
	}
}
