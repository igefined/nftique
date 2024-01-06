package middleware

import "github.com/gofiber/fiber/v2"

const (
	realIPHeader    = "X-Real-Ip"
	userAgentHeader = "User-Agent"
)

func GetRealIPFromContext(ctx *fiber.Ctx) string {
	xRealIPs := ctx.GetReqHeaders()[realIPHeader]
	if len(xRealIPs) > 0 {
		return xRealIPs[0]
	}
	return ""
}

func GetUserAgentFromContext(ctx *fiber.Ctx) string {
	userAgents := ctx.GetReqHeaders()[userAgentHeader]
	if len(userAgents) > 0 {
		return userAgents[0]
	}
	return ""
}
