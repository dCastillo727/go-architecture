package driven

import (
	"context"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
)

type ExampleRepositoryPort interface {
	FindAll(ctx context.Context) ([]example.Example, error)
}
