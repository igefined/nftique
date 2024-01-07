package middleware

import (
	"strings"

	apiErr "github.com/igefined/nftique/internal/transport/http/error"
	"github.com/igefined/nftique/pkg/rate_limiter"

	"github.com/gofiber/fiber/v2"
)

func RateLimit(limiter *rate_limiter.RateLimiter) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		identifier, requestType, err := GetClientIdentifierAndRequestType(ctx)
		if err != nil {
			return apiErr.Respond(ctx, err)
		}

		isAllowed := limiter.IsAllowed(ctx.Context(), requestType, identifier)
		if !isAllowed {
			return apiErr.Respond(ctx, apiErr.ErrTooManyRequests)
		}

		return ctx.Next()
	}
}

func GetClientIdentifierAndRequestType(ctx *fiber.Ctx) (string, rate_limiter.RequestType, error) {
	ip := ctx.IP()
	if ip == "" {
		return "",
			rate_limiter.Undefined,
			apiErr.ErrXRealIPRequired
	}

	url := ctx.Path()
	if strings.Contains(url, "/auth") {
		return ip, rate_limiter.Auth, nil
	}

	return ip, rate_limiter.Common, nil
}
