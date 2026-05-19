package api_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Register(rg *gin.RouterGroup) {
	group := rg.Group("/ping")
	{
		group.GET("/", h.pong)
	}
}

func (h *PingHandler) pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
