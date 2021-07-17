//+build wireinject

package main

import (
	"infra-practice/pkg/presentation/controller"

	"github.com/google/wire"
)

var ProvideApp = wire.NewSet(
	ProvidePresentation,
	NewServer,
)

var ProvidePresentation = wire.NewSet(
	controller.NewSampleController,
)

func initialize() (*Server, error) {
	wire.Build(ProvideApp)
	return nil, nil
}
