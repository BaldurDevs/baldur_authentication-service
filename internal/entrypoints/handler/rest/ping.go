package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandlerFactory() PingHandler {
	return NewPingHandler()
}

func NewPingHandler() PingHandler {
	return PingHandler{}
}

type PingHandler struct{}

func (handler *PingHandler) RegisterRoutes(r *gin.Engine) {
	r.GET(basePath+"/ping", handler.getPing)
}

func (handler *PingHandler) getPing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
