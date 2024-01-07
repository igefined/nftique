package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LoggerRequest(logger *zap.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			logger.Error("request",
				zap.String("path", ctx.Path()),
				zap.String("ip", ctx.IP()),
				zap.Error(err),
			)

			return err
		}

		logger.Info("request",
			zap.String("path", ctx.Path()),
			zap.String("ip", ctx.IP()),
			zap.Int("status", ctx.Response().StatusCode()),
		)

		return nil
	}
}
