package handler

import (
	"context"

	"github.com/igefined/nftique/internal/rate_limiter"
	"github.com/igefined/nftique/internal/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	Services,
	rate_limiter.Module,
	fx.Provide(
		NewNFTHandler,
		NewAuthHandler,
		NewHandler,
	),
	fx.Invoke(func(ls fx.Lifecycle) {
		ls.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				return nil
			},
			OnStop: func(_ context.Context) error {
				return nil
			},
		})
	}),
)

var Services = fx.Options(
	fx.Provide(
		fx.Annotate(service.NewNFTService, fx.As(new(NFTService))),
	),
)
