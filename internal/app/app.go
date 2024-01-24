package app

import (
	"context"
	"fmt"

	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/internal/transport/http"
	"github.com/igefined/nftique/internal/transport/http/handler"
	"github.com/igefined/nftique/pkg/sys"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const Name = "nftique"

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
	sys  *sys.Sys
	http *http.App
}

func NewWebServer(cfg *config.Config, logger *zap.Logger, httpServer *http.App) *WebServer {
	inf := sys.NewAppInfo(Name).
		WithVersion(Version).WithBuildCommit(Commit).WithBuildTime(BuildDate)
	sysServer := sys.NewSys(logger, &cfg.MainCfg, inf)

	instance := &WebServer{
		cfg:    cfg,
		logger: logger,
		http:   httpServer,
		sys:    sysServer,
		app:    fiber.New(),
	}

	return instance
}

func (s *WebServer) StartServer() error {
	s.app.Use(s.sys.PrometheusMiddleware())
	s.app.Mount("/api", s.http.App)

	s.sys.Start()

	return s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port))
}

func (s *WebServer) ShutDownServer(ctx context.Context) error {
	if err := s.sys.Stop(ctx); err != nil {
		return err
	}

	return s.app.Shutdown()
}
