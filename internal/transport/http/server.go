package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/igefined/nftique/internal/transport/http/handler"
)

type App struct {
	handler handler.Handler

	*fiber.App
}

func NewServer(handler handler.Handler) *App {
	instance := &App{
		App:     fiber.New(),
		handler: handler,
	}

	instance.initRoutes()

	return instance
}

func (a *App) initRoutes() {
	v1 := a.Group("/v1")
	v1.Get("/nft", a.handler.NFTHandler().ListAllNFTs)
}
