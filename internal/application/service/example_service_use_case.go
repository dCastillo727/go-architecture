package service

import (
	"context"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
	"github.com/dCastillo727/go-architecture/internal/application/port/driven"
	"github.com/dCastillo727/go-architecture/internal/registry"
)

type ExampleServiceUseCase struct {
	repository driven.ExampleRepositoryPort
}

func NewExampleServiceUseCase(repository *registry.Repository) *ExampleServiceUseCase {
	return &ExampleServiceUseCase{
		repository: repository.Example,
	}
}

func (s *ExampleServiceUseCase) GetAll(ctx context.Context) ([]example.Example, error) {
	return s.repository.FindAll(ctx)
}

func (s *ExampleServiceUseCase) GetByID(ctx context.Context, id int64) (*example.Example, error) {
	return s.repository.FindByID(ctx, id)
}
