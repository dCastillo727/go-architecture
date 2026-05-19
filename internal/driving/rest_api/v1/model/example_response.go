package api_model

import "github.com/dCastillo727/go-architecture/internal/application/domain/example"

type ExampleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewExampleResponse(model example.Example) *ExampleResponse {
	return &ExampleResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}

func NewExampleResponses(models []example.Example) []*ExampleResponse {
	responses := make([]*ExampleResponse, len(models))
	for i, model := range models {
		responses[i] = NewExampleResponse(model)
	}
	return responses
}
