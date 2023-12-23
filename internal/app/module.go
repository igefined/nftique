package app

import (
	"context"

	"github.com/igefined/nftique/internal/config"
	cfg "github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/log"
	"github.com/igefined/nftique/pkg/sys"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		func() context.Context {
			return cfg.SigTermIntCtx()
		},
	),
	log.Module,
	sys.Module,
	config.Module,
	WebServerModule,
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(
			zap.String("version", Version),
			zap.String("commit", Commit),
			zap.String("buildDate", BuildDate),
		)
	}),
	fx.Invoke(
		func(ls fx.Lifecycle,
			appCtx context.Context,
			logger *zap.Logger,
			cfg *config.Config,
			webServer *WebServer,
		) {
			ls.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					go func() {
						if err := webServer.StartServer(); err != nil {
							logger.Fatal("start server error", zap.Error(err))
						}
					}()

					return nil
				},
				OnStop: func(_ context.Context) error {
					return webServer.ShutDownServer()
				},
			})
		},
	),
)
