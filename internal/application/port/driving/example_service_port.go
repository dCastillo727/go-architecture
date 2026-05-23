package driving

import (
	"context"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
)

type ExampleService interface {
	GetAll(ctx context.Context) ([]example.Example, error)
	GetByID(ctx context.Context, id int64) (*example.Example, error)
}
