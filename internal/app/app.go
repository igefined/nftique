package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/pkg/sys"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// ldflags
	Commit    string
	BuildDate string
	Version   string
)

var WebServerModule = fx.Options(fx.Provide(NewWebServer))

type WebServer struct {
	cfg    *config.Config
	logger *zap.Logger

	// apps
	server *fiber.App
	sys    *sys.App
}

func NewWebServer(cfg *config.Config, logger *zap.Logger, sys *sys.App) *WebServer {
	return &WebServer{cfg: cfg, logger: logger, sys: sys, server: fiber.New()}
}

func (s *WebServer) StartServer() error {
	prometheus := sys.NewPrometheus(s.cfg.Namespace, s.cfg.Environment, s.cfg.ServiceName)
	prometheus.RegisterAt(s.server, "/system/metrics")
	s.server.Use(prometheus.Middleware)

	s.server.Mount("/system", s.sys.App)

	return s.server.Listen(fmt.Sprintf(":%s", s.cfg.Port))
}

func (s *WebServer) ShutDownServer() error {
	return s.server.Shutdown()
}
