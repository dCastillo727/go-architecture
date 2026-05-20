package restapi

import (
	"net/http"

	"github.com/dCastillo727/go-architecture/internal/application/port/driving"
	restapi "github.com/dCastillo727/go-architecture/internal/driving/rest_api/v1/model"
	"github.com/gin-gonic/gin"
)

type exampleHandler struct {
	service driving.ExampleService
}

func NewExampleHandler(service driving.ExampleService) *exampleHandler {
	return &exampleHandler{
		service: service,
	}
}

func (h *exampleHandler) Register(rg *gin.RouterGroup) {
	group := rg.Group("/examples")
	{
		group.GET("/", h.findAll)
	}
}

func (h *exampleHandler) findAll(c *gin.Context) {
	ctx := c.Request.Context()

	examples, err := h.service.FindAll(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restapi.NewExampleResponses(examples))
}
