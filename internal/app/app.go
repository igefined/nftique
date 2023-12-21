package app

import (
	"github.com/igefined/nftique/internal/config"
	pkgConfig "github.com/igefined/nftique/pkg/config"
)

var (
	// ldflags
	Commit    string
	BuildDate string
	Version   string
)

func Run() {
	ctx := pkgConfig.SigTermIntCtx()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	_ = cfg
	_ = ctx
}
