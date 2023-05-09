package bootstrap

import (
	"context"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/infrastructure/handler"
	"github.com/joho/godotenv"
)

func Run(boot []byte) {
	_ = godotenv.Load()

	db := newDatabase()

	ginEntry := newGinEntry(boot)
	ginEntry.Bootstrap(context.Background())

	api := ginEntry.Router

	handler.InitRoutes(model.RouterSpecification{
		Api: api,
		DB:  db,
	})

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(context.Background())
}
