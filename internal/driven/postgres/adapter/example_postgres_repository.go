package postgres

import (
	"context"
	"errors"

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

func (r *ExamplePostgresRepository) FindByID(ctx context.Context, id int64) (*example.Example, error) {
	query := "SELECT id, name FROM example WHERE id = $1"

	row := r.pool.QueryRow(ctx, query, id)

	var dbo postgres.ExampleDBO
	if err := row.Scan(&dbo.ID, &dbo.Name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, example.ErrExampleNotFound
		}
		return nil, err
	}

	return dbo.ToModel(), nil
}

func (r *ExamplePostgresRepository) Create(ctx context.Context, e *example.Example) (*example.Example, error) {
	query := "INSERT INTO example (name) VALUES ($1) RETURNING id"

	err := r.pool.QueryRow(ctx, query, e.Name).Scan(&e.ID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *ExamplePostgresRepository) Update(ctx context.Context, e *example.Example) (*example.Example, error) {
	query := "UPDATE example SET name = $1 WHERE id = $2"

	cmdTag, err := r.pool.Exec(ctx, query, e.Name, e.ID)
	if err != nil {
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, example.ErrExampleNotFound
	}

	return e, nil
}
