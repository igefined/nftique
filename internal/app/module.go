package app

import (
	"context"

	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/internal/rate_limiter"
	cfg "github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/log"
	"github.com/igefined/nftique/pkg/sys"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(func() context.Context { return cfg.SigTermIntCtx() }), //nolint:gocritic
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
					settings := rate_limiter.NewSettings()
					if err := settings.Set(rate_limiter.Auth, 1, 5); err != nil {
						return err
					}

					if err := settings.Set(rate_limiter.Common, 1, 5); err != nil {
						return err
					}

					rate_limiter.NewRateLimiter(logger, settings, nil)

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
