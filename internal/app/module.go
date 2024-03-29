package app

import (
	"context"

	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/internal/migrations"
	cfg "github.com/igefined/nftique/pkg/config"
	db "github.com/igefined/nftique/pkg/db/postgresql"
	"github.com/igefined/nftique/pkg/log"
	"github.com/igefined/nftique/pkg/validator"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(func() context.Context { return cfg.SigTermIntCtx() }), //nolint:gocritic
	log.Module,
	config.Module,
	validator.Module,
	WebServerModule,
	fx.Provide(db.New),
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
			_ *db.QBuilder,
			webServer *WebServer,
		) {
			ls.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					if err := db.Migrate(logger, &migrations.DB, &cfg.DBCfg); err != nil {
						return err
					}

					go func() {
						if err := webServer.StartServer(); err != nil {
							logger.Fatal("start server error", zap.Error(err))
						}
					}()

					return nil
				},
				OnStop: func(ctx context.Context) error {
					return webServer.ShutDownServer(ctx)
				},
			})
		},
	),
)
