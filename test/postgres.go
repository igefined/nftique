package test

import (
	"context"
	"time"

	cfg "github.com/igefined/nftique/pkg/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	cfg *cfg.DBCfg
	*postgres.PostgresContainer
}

func NewPostgresContainer(ctx context.Context, cfg *cfg.DBCfg, opt *Opt) (*PostgresContainer, error) {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage(opt.Image),
		postgres.WithDatabase(cfg.GetDatabaseName()),
		postgres.WithUsername(cfg.GetDatabaseUser()),
		postgres.WithPassword(cfg.GetDatabasePassword()),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{cfg: cfg, PostgresContainer: container}, nil
}
