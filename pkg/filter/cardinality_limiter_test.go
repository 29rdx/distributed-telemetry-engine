package filter

import (
	"testing"
)

func TestCardinalityLimiter(t *testing.T) {
	limiter := NewCardinalityLimiter(2)
	if !limiter.Allow("env", "prod") {
		t.Fatal("expected first tag to be allowed")
	}
	if !limiter.Allow("env", "staging") {
		t.Fatal("expected second tag to be allowed")
	}
	// Third unique tag exceeds maxKeys=2
	if limiter.Allow("env", "dev") {
		t.Fatal("expected third tag to be rejected by cardinality limit")
	}
}
