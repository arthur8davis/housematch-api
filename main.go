package main

import (
	"embed"
	"github.com/Melany751/house-match-server/bootstrap"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

//go:embed config/docs
var docsFS embed.FS

func init() {
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.SWEntryType, "users", &docsFS)
}

//go:embed boot.yaml
var boot []byte

func main() {
	bootstrap.Run(boot)
}
