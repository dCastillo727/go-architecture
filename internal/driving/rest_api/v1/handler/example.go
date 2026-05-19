package api_handler

import (
	"net/http"

	"github.com/dCastillo727/go-architecture/internal/application/port/driving"
	api_model "github.com/dCastillo727/go-architecture/internal/driving/rest_api/v1/model"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	service driving.ExampleService
}

func NewExampleHandler(service driving.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		service: service,
	}
}

func (h *ExampleHandler) Register(rg *gin.RouterGroup) {
	group := rg.Group("/examples")
	{
		group.GET("/", h.findAll)
	}
}

func (h *ExampleHandler) findAll(c *gin.Context) {
	examples, err := h.service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api_model.NewExampleResponses(examples))
}
