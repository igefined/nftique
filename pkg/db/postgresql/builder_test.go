package postgresql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseName(t *testing.T) {
	const (
		url    = "postgres://postgres:postgres@localhost:5466/test_clients?sslmode=disable"
		expect = "test_clients"
	)

	assert.Equal(t, GetDatabaseName(url), expect)
}

func TestReplaceDbName(t *testing.T) {
	tCases := []struct {
		srcUrl    string
		dbName    string
		resultUrl string
	}{
		{
			srcUrl:    "postgres://postgres:postgres@localhost:5466/test?sslmode=disable",
			dbName:    "silly",
			resultUrl: "postgres://postgres:postgres@localhost:5466/silly?sslmode=disable",
		},
		{
			srcUrl:    "postgres://postgres:postgres@localhost:5466/test",
			dbName:    "silly",
			resultUrl: "postgres://postgres:postgres@localhost:5466/silly",
		},
		{
			srcUrl:    "postgres://postgres:12345@localhost:5432/common?sslmode=disable&pool_max_conns=16&pool_max_conn_idle_time=30m&pool_max_conn_lifetime=1h&pool_health_check_period=1m", //nolint:lll
			dbName:    "nh_common",
			resultUrl: "postgres://postgres:12345@localhost:5432/nh_common?sslmode=disable&pool_max_conns=16&pool_max_conn_idle_time=30m&pool_max_conn_lifetime=1h&pool_health_check_period=1m", //nolint:lll
		},
	}

	for _, c := range tCases {
		assert.Equal(t, ReplaceDbName(c.srcUrl, c.dbName), c.resultUrl)
	}
}

//nolint:gocritic
//func TestCreateDatabase(t *testing.T) {
//	var isExists bool
//	logger, err := log.NewLogger(zap.DebugLevel)
//	assert.NoError(t, err)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
//	defer cancel()
//
//	type testConfig struct {
//		sync.RWMutex
//		config.DBCfg `mapstructure:",squash"`
//	}
//
//	var cfg *testConfig
//	assert.NoError(t, config.GetConfig(&cfg, []*config.EnvVar{}))
//
//	assert.NotEmpty(t, cfg.URL)
//
//	CreateDatabase(ctx, logger, cfg.URL)
//
//	pool, err := pgxpool.New(ctx, ReplaceDbName(cfg.URL, "postgres"))
//	assert.NoError(t, err)
//
//	checkingSql := `select exists(select datname from pg_catalog.pg_database where datname = $1) as exist`
//	row := pool.QueryRow(ctx, checkingSql, GetDatabaseName(cfg.URL))
//	err = row.Scan(&isExists)
//	assert.NoError(t, err)
//	assert.True(t, isExists)
//
//	dropDbSql := fmt.Sprintf("drop database if exists %s", GetDatabaseName(cfg.URL))
//	_, err = pool.Query(ctx, dropDbSql)
//	assert.NoError(t, err)
//}
