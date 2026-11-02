package filter

import (
	"sync"
)

type CardinalityLimiter struct {
	mu        sync.RWMutex
	maxKeys   int
	seenKeys  map[string]struct{}
}

func NewCardinalityLimiter(maxKeys int) *CardinalityLimiter {
	return &CardinalityLimiter{
		maxKeys:  maxKeys,
		seenKeys: make(map[string]struct{}),
	}
}

func (c *CardinalityLimiter) Allow(tagKey, tagValue string) bool {
	combined := tagKey + "=" + tagValue
	c.mu.RLock()
	_, exists := c.seenKeys[combined]
	c.mu.RUnlock()
	if exists {
		return true
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.seenKeys) >= c.maxKeys {
		return false // Tag cardinality budget exhausted
	}
	c.seenKeys[combined] = struct{}{}
	return true
}
