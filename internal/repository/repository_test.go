package repository

import (
	"context"
	"sync"
	"testing"

	schema "github.com/igefined/nftique/internal/migrations"
	"github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/db/postgresql"

	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

var (
	once   sync.Once
	qb     *postgresql.QBuilder
	logger *zap.Logger
)

func TestMain(m *testing.M) {
	once.Do(func() {
		var err error

		loggerCfg := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger, err = loggerCfg.Build()
		if err != nil {
			log.Fatal(err)
		}

		type testConfig struct {
			sync.RWMutex
			config.DBCfg `mapstructure:",squash"`
		}

		var cfg *testConfig
		if err = config.GetConfig(&cfg, []*config.EnvVar{}); err != nil {
			log.Fatal(err)
		}

		if err = postgresql.Migrate(logger, &schema.DB, &cfg.DBCfg); err != nil {
			log.Fatal(err)
		}

		qb = postgresql.New(logger, &cfg.DBCfg, nil)
	})

	m.Run()
}

func deleteRows(ctx context.Context, table string) (err error) {
	_, err = qb.Querier().Exec(ctx, "delete from "+table)

	return
}

func clearDB(t *testing.T) {
	if err := deleteRows(context.Background(), "users"); err != nil {
		t.Fatalf("cannot clear users: %v", err)
	}
}
