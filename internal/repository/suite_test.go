package repository

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/igefined/nftique/internal/domain"
	schema "github.com/igefined/nftique/internal/migrations"
	"github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/db/postgresql"
	"github.com/igefined/nftique/pkg/log"
	"github.com/igefined/nftique/test"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

const defaultPostgresImage = "postgres:15.3-alpine"

type Suite struct {
	suite.Suite
	ctx context.Context

	cfg        *config.DBCfg
	logger     *zap.Logger
	container  *test.PostgresContainer
	qb         *postgresql.QBuilder
	repository *UserRepository
}

func TestBuilderSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	logger, err := log.NewLogger(zap.DebugLevel)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	_ = cancel

	type testConfig struct {
		sync.RWMutex
		config.DBCfg `mapstructure:",squash"`
	}

	var cfg *testConfig
	s.Require().NoError(config.GetConfig(&cfg, []*config.EnvVar{}))
	s.Require().NotEmpty(cfg.URL)

	s.logger = logger
	s.cfg = &cfg.DBCfg
	s.ctx = ctx

	pgContainer, err := test.NewPostgresContainer(ctx, s.cfg, &test.Opt{Enabled: true, Image: defaultPostgresImage})
	s.Require().NoError(err)

	s.container = pgContainer

	if err = postgresql.Migrate(logger, &schema.DB, &cfg.DBCfg); err != nil {
		s.logger.Fatal("postgresql migrate error", zap.Error(err))
	}

	s.qb = postgresql.New(logger, &cfg.DBCfg, nil)

	s.repository = NewUserRepository(s.qb)
}

func (s *Suite) TearDownSuite() {
	if err := s.container.Terminate(s.ctx); err != nil {
		s.logger.Error("error terminating postgres container", zap.Error(err))
	}
}

func (s *Suite) deleteRows(ctx context.Context, table string) (err error) {
	_, err = s.qb.Querier().Exec(ctx, "delete from "+table)

	return
}

func (s *Suite) clearDB() {
	if err := s.deleteRows(context.Background(), "users"); err != nil {
		s.logger.Fatal("cannot clear users", zap.Error(err))
	}
}

func (s *Suite) createUser(ctx context.Context) (*domain.User, error) {
	usr := fakeUser()
	q := `insert into users(id, web3_address, first_name, last_name, username) values ($1, $2, $3, $4, $5)`
	_, err := s.qb.Querier().Exec(ctx, q, usr.UID, usr.Web3Address, usr.FirstName, usr.LastName, usr.Username)
	if err != nil {
		return nil, errors.New("error inserted fake user")
	}

	return usr, err
}

func fakeUser() *domain.User {
	return &domain.User{
		UID:         uuid.New(),
		FirstName:   randomdata.FirstName(randomdata.RandomGender),
		LastName:    randomdata.LastName(),
		Username:    randomdata.PhoneNumber(),
		Web3Address: randomdata.Address(),
	}
}
