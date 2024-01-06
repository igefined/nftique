package rate_limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	sync.Mutex
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

func NewTokenBucket(settings *settings) *TokenBucket {
	return &TokenBucket{
		tokens:         settings.MaxTokens,
		maxTokens:      settings.MaxTokens,
		refillRate:     settings.Rate,
		lastRefillTime: time.Now(),
	}
}

func (b *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(b.lastRefillTime)
	tokensToAdd := b.refillRate * duration.Seconds()
	b.tokens = tokensToAdd
	b.lastRefillTime = now
}

func (b *TokenBucket) IsAllowed(tokens float64) bool {
	b.Lock()
	defer b.Unlock()

	b.refill()
	if b.tokens >= tokens {
		b.tokens -= tokens
		return true
	}

	return false
}
