package postgres

import (
	"context"
	"fmt"

	postgres "github.com/dCastillo727/go-architecture/internal/driven/postgres/adapter"
	"github.com/dCastillo727/go-architecture/internal/registry"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(ctx context.Context, config registry.Config) (*PostgresClient, *registry.Repository, error) {
	pool, err := connectPostgres(ctx, config)
	if err != nil {
		return nil, nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, nil, err
	}

	return &PostgresClient{
			pool: pool,
		}, &registry.Repository{
			Example: postgres.NewExamplePostgresRepository(pool),
		}, nil
}

func (c *PostgresClient) Close() error {
	c.pool.Close()
	return nil
}

func connectPostgres(ctx context.Context, config registry.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DdbbUser, config.DdbbPassword, config.DdbbHost, config.DdbbName)

	return pgxpool.New(ctx, connStr)
}
