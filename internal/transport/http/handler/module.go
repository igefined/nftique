package handler

import (
	"github.com/igefined/nftique/internal/service"
	"github.com/igefined/nftique/pkg/rate_limiter"

	"go.uber.org/fx"
)

var Module = fx.Options(
	Services,
	service.Module,
	rate_limiter.Module,
	fx.Provide(
		NewNFTHandler,
		NewHandler,
	),
)

var Services = fx.Options(
	fx.Provide(
		fx.Annotate(service.NewNFTService, fx.As(new(NFTService))),
	),
)
