package service

import (
	"github.com/igefined/nftique/internal/repository"
	"github.com/igefined/nftique/pkg/s3"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewNFTService,
		s3.New,
		fx.Annotate(repository.NewUserRepository, fx.As(new(UserRepository))),
	),
)
