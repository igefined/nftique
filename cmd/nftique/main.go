package main

import (
	"go.uber.org/fx"

	"github.com/igefined/nftique/internal/app"
)

func main() {
	fx.New(app.Module).Run()
}
