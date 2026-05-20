package postgres

import (
	"context"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
	postgres "github.com/dCastillo727/go-architecture/internal/driven/postgres/dbo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExamplePostgresRepository struct {
	pool *pgxpool.Pool
}

func NewExamplePostgresRepository(pool *pgxpool.Pool) *ExamplePostgresRepository {
	return &ExamplePostgresRepository{
		pool: pool,
	}
}

func (r *ExamplePostgresRepository) FindAll(ctx context.Context) ([]example.Example, error) {
	query := "SELECT id, name FROM example"

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	examples, err := pgx.CollectRows(rows, pgx.RowToStructByName[postgres.ExampleDBO])
	if err != nil {
		return nil, err
	}

	return postgres.ExampleDBOsToModels(examples), nil
}
