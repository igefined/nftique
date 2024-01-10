package migrations

import "embed"

//go:embed files/*.sql
var DB embed.FS
