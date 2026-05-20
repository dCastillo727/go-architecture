package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingHandler struct {
}

func NewPingHandler() *pingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Register(rg *gin.RouterGroup) {
	group := rg.Group("/ping")
	{
		group.GET("/", h.pong)
	}
}

func (h *pingHandler) pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
