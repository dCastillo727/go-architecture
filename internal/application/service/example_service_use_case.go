package service

import "github.com/dCastillo727/go-architecture/internal/application/domain/example"

type ExampleServiceUseCase struct {
}

func NewExampleServiceUseCase() *ExampleServiceUseCase {
	return &ExampleServiceUseCase{}
}

func (s *ExampleServiceUseCase) FindAll() ([]example.Example, error) {
	return []example.Example{
		{
			ID:   1,
			Name: "example1",
		},
		{
			ID:   2,
			Name: "example2",
		},
	}, nil
}
