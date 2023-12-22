package app

import (
	"context"

	"github.com/igefined/nftique/internal/config"
	cfg "github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/log"

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
	config.Module,
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(
			zap.String("version", Version),
			zap.String("commit", Commit),
			zap.String("buildDate", BuildDate),
		)
	}),
	fx.Invoke(
		func(ls fx.Lifecycle, cxt context.Context, logger *zap.Logger, cfg *config.Config) {
			ls.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return Run()
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		},
	),
)
