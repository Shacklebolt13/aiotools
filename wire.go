//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeApp() Application {
	wire.Build(
		NewApplication,
		NewServer,
		NewShortenerServiceHandler,
		NewDBConnection,
		wire.Bind(new(Application), new(*ApplicationImpl)),
	)
	return &ApplicationImpl{}
}
