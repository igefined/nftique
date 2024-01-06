package middleware

import (
	"strings"

	apiErr "github.com/igefined/nftique/internal/transport/http/error"

	"github.com/gofiber/fiber/v2"
	"github.com/igefined/nftique/internal/rate_limiter"
)

func RateLimiter(ctx *fiber.Ctx) error {
	identifier, requestType, err := GetClientIdentifierAndRequestType(ctx)
	if err != nil {
		return apiErr.Respond(ctx, err)
	}

	limiter := rate_limiter.GetRateLimiter()
	if err != nil {
		return apiErr.Respond(ctx, apiErr.ErrInternal)
	}

	isAllowed := limiter.IsAllowed(requestType, 5, identifier)
	if !isAllowed {
		return apiErr.Respond(ctx, apiErr.ErrTooManyRequests)
	}

	return ctx.Next()
}

func GetClientIdentifierAndRequestType(ctx *fiber.Ctx) (string, rate_limiter.RequestType, error) {
	ip := GetRealIPFromContext(ctx)
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
