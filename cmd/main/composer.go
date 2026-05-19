package main

import (
	"github.com/dCastillo727/go-architecture/internal/application/service"
	"github.com/dCastillo727/go-architecture/internal/registry"
)

func composeApp() *registry.App {
	services := registry.Services{
		Example: service.NewExampleServiceUseCase(),
	}

	return registry.NewApp(services)
}
