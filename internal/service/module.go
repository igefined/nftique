package service

import (
	"github.com/igefined/nftique/internal/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewNFTService,
		fx.Annotate(repository.NewUserRepository, fx.As(new(UserRepository))),
	),
)
