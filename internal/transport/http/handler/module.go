package handler

import (
	"context"

	"github.com/igefined/nftique/internal/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	Services,
	fx.Provide(
		NewNFTHandler,
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
