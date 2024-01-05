package rate_limiter

import "sync"

var (
	once    sync.Once
	limiter *Limiter
)

type Limiter struct {
	settings *Settings
}

func NewRateLimiter(settings *Settings) *Limiter {
	once.Do(func() {
		limiter = &Limiter{
			settings: settings,
		}
	})

	return limiter
}
