package postgres

import "github.com/dCastillo727/go-architecture/internal/application/domain/example"

type ExampleDBO struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (d *ExampleDBO) ToModel() *example.Example {
	return &example.Example{
		ID:   d.ID,
		Name: d.Name,
	}
}

func ExampleDBOsToModels(dbos []ExampleDBO) []example.Example {
	models := make([]example.Example, len(dbos))
	for i, dbo := range dbos {
		models[i] = *dbo.ToModel()
	}
	return models
}
