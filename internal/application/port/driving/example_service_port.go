package driving

import (
	"context"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
)

type ExampleService interface {
	FindAll(ctx context.Context) ([]example.Example, error)
}
