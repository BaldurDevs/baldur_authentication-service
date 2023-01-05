package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandlerFactory() *PingHandler {
	return NewPingHandler()
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

type PingHandler struct{}

func (handler *PingHandler) RegisterRouters(r *gin.Engine) {
	r.GET(basePath+"/ping", handler.ping)
}

func (handler *PingHandler) ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}