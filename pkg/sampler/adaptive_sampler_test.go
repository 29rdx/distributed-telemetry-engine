package sampler

import (
	"testing"
)

func TestAdaptiveSamplerErrors(t *testing.T) {
	s := NewAdaptiveSampler(0.0) // 0% standard sampling
	if !s.ShouldSample(true) {
		t.Fatal("errors must always be sampled under errorBoost policy")
	}
	if s.ShouldSample(false) {
		t.Fatal("clean traces must not be sampled when sampleRate is 0.0")
	}
}
