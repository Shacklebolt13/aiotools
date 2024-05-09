//go:build wireinject
// +build wireinject

package main

import (
	"aiotools/server/internal/database"
	"aiotools/server/internal/database/model"
	"aiotools/server/internal/handlers"
	"aiotools/server/internal/services"

	"github.com/google/wire"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDialector(conf AppConfig) gorm.Dialector {
	return postgres.Open(conf.DSN)
}
func InitializeApp(severOpts []grpc.ServerOption, gormOpts []gorm.Option, conf AppConfig) (ApplicationImpl, error) {
	wire.Build(
		//application (to be injected)
		NewApplication,

		//handlers
		handlers.NewShortenerServiceHandler,
		handlers.NewPubSubServiceHandler,

		//services
		services.NewShortenService,
		services.NewTopicService,

		//repositories
		model.NewURLBaseRepository,

		//db and gorm
		database.NewDatabase,
		gorm.Open,
		newDialector,

		//grpc
		grpc.NewServer,
	)
	return ApplicationImpl{}, nil
}
