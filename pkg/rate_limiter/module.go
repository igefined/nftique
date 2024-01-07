package rate_limiter

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewRateLimiter,
		NewSettings,
	),
)
