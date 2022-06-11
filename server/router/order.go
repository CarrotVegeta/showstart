package router

import (
	. "github.com/CarrotVegeta/showstart/controller"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/gin-gonic/gin"
)

type OrderServer struct {
}

func (o *OrderServer) RegisterRouter(e *gin.Engine) {
	e.POST("/order", server.Handler()(Order))
}
