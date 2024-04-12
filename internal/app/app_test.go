package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/igefined/nftique/internal/config"
	"github.com/igefined/nftique/internal/transport/http"
	"github.com/igefined/nftique/pkg/log"
)

func TestNewWebServer(t *testing.T) {
	cfg, err := config.New()
	assert.NoError(t, err)

	lg, err := log.NewLogger(zap.DebugLevel)
	assert.NoError(t, err)

	server := NewWebServer(cfg, lg, &http.App{})
	assert.NotNil(t, server)
}
