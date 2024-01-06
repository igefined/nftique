package http

import (
	"github.com/igefined/nftique/internal/transport/http/handler"
	"github.com/igefined/nftique/internal/transport/http/middleware"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	logger  *zap.Logger
	handler handler.Handler

	*fiber.App
}

func NewServer(logger *zap.Logger, handler handler.Handler) *App {
	instance := &App{
		logger:  logger,
		App:     fiber.New(),
		handler: handler,
	}

	instance.initRoutes()

	return instance
}

func (a *App) initRoutes() {
	v1 := a.Group("/v1")
	v1.Use(middleware.RateLimiter)

	authV1 := v1.Group("/auth")
	authV1.Post("/sign_in", a.handler.AuthHandler().SignIn)

	nftV1 := v1.Group("/nfts")
	nftV1.Get("/", a.handler.NFTHandler().ListAllNFTs)
}
