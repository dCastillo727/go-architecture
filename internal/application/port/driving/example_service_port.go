package driving

import "github.com/dCastillo727/go-architecture/internal/application/domain/example"

type ExampleService interface {
	FindAll() ([]example.Example, error)
}
