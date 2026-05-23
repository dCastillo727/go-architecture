package restapi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dCastillo727/go-architecture/internal/application/domain/example"
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
		group.GET("/", h.getAll)
		group.GET("/:id", h.getByID)
	}
}

func (h *exampleHandler) getAll(c *gin.Context) {
	ctx := c.Request.Context()

	examples, err := h.service.GetAll(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restapi.NewExampleResponses(examples))
}

func (h *exampleHandler) getByID(c *gin.Context) {
	rawId := c.Param("id")
	id, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx := c.Request.Context()

	result, err := h.service.GetByID(ctx, id)

	if err != nil {
		var code int
		switch {
		case errors.Is(err, example.ErrExampleNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restapi.NewExampleResponse(*result))
}
