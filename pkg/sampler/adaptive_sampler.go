package sampler

import (
	"math/rand"
	"sync"
	"time"
)

type AdaptiveSampler struct {
	mu           sync.RWMutex
	sampleRate   float64
	errorBoost   bool
	rng          *rand.Rand
}

func NewAdaptiveSampler(initialRate float64) *AdaptiveSampler {
	return &AdaptiveSampler{
		sampleRate: initialRate,
		errorBoost: true,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *AdaptiveSampler) ShouldSample(isError bool) bool {
	if isError && s.errorBoost {
		return true // Always sample errors
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rng.Float64() <= s.sampleRate
}
