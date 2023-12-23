package sys

import (
	"encoding/json"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	*fiber.App
}

func NewApp(logger *zap.Logger) *App {
	sys := fiber.New()
	sys.Use(fiberzap.New(fiberzap.Config{Logger: logger}))
	sys.Get("/healthcheck", HealthcheckHandler)

	return &App{App: sys}
}

func HealthcheckHandler(ctx *fiber.Ctx) error {
	type resp struct {
		Status string `json:"status"`
	}

	body, err := json.Marshal(&resp{Status: "ok"})
	if err != nil {
		return err
	}

	return ctx.Send(body)
}
