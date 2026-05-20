package main

import (
	"context"
	"log"

	"github.com/dCastillo727/go-architecture/internal/application/service"
	"github.com/dCastillo727/go-architecture/internal/driven/postgres"
	"github.com/dCastillo727/go-architecture/internal/registry"
)

type ApplicationContext struct {
	Services registry.Services
	DBClient *postgres.PostgresClient
}

func (a *ApplicationContext) Close() error {
	return a.DBClient.Close()
}

func composeApp(cfg *registry.Config) (*ApplicationContext, error) {
	ctx := context.Background()

	pgClient, pgRepos, err := postgres.NewPostgresRepository(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	services := composeServices(pgRepos)

	return &ApplicationContext{
		Services: services,
		DBClient: pgClient,
	}, nil
}

func composeServices(r *registry.Repository) registry.Services {
	return registry.Services{
		Example: service.NewExampleServiceUseCase(r),
	}
}
