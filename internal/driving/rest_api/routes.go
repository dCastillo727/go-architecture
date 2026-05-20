package restapi

import (
	restapi "github.com/dCastillo727/go-architecture/internal/driving/rest_api/v1/handler"
	"github.com/dCastillo727/go-architecture/internal/registry"
	"github.com/gin-gonic/gin"
)

type handler interface {
	Register(rg *gin.RouterGroup)
}

func SetupRouter(s registry.Services) *gin.Engine {
	router := gin.Default()
	handlers := setupHandlers(s)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	for _, handler := range handlers {
		handler.Register(v1)
	}

	return router
}

func setupHandlers(s registry.Services) []handler {
	return []handler{
		restapi.NewExampleHandler(s.Example),
		restapi.NewPingHandler(),
	}
}
