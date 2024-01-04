package app

import (
	"fmt"

	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/internal/transport/http"
	"github.com/igefined/nftique/internal/transport/http/handler"
	"github.com/igefined/nftique/pkg/sys"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// ldflags
	Commit    string
	BuildDate string
	Version   string
)

var WebServerModule = fx.Options(
	handler.Module,
	fx.Provide(
		NewWebServer,
		http.NewServer,
	),
)

type WebServer struct {
	cfg    *config.Config
	logger *zap.Logger

	// apps
	app  *fiber.App
	http *http.App
	sys  *sys.App
}

func NewWebServer(cfg *config.Config, logger *zap.Logger, sys *sys.App, httpServer *http.App) *WebServer {
	return &WebServer{cfg: cfg, logger: logger, sys: sys, http: httpServer, app: fiber.New()}
}

func (s *WebServer) StartServer() error {
	prometheus := sys.NewPrometheus(s.cfg.Namespace, s.cfg.Environment, s.cfg.ServiceName)
	prometheus.RegisterAt(s.app, "/system/metrics")
	s.app.Use(prometheus.Middleware)

	s.app.Mount("/system", s.sys.App)
	s.app.Mount("/api", s.http.App)

	return s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port))
}

func (s *WebServer) ShutDownServer() error {
	return s.app.Shutdown()
}
