package aggregator

import (
	"sort"
	"sync"
)

type QuantileTracker struct {
	mu      sync.Mutex
	samples []float64
	maxCap  int
}

func NewQuantileTracker(capacity int) *QuantileTracker {
	return &QuantileTracker{
		samples: make([]float64, 0, capacity),
		maxCap:  capacity,
	}
}

func (q *QuantileTracker) Record(val float64) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.samples) < q.maxCap {
		q.samples = append(q.samples, val)
	}
}

func (q *QuantileTracker) GetPercentile(p float64) float64 {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.samples) == 0 {
		return 0.0
	}
	sorted := make([]float64, len(q.samples))
	copy(sorted, q.samples)
	sort.Float64s(sorted)
	idx := int(p * float64(len(sorted)-1))
	return sorted[idx]
}
