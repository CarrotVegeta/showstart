package router

import (
	"github.com/CarrotVegeta/showstart/api"
	. "github.com/CarrotVegeta/showstart/pkg/service"
	"github.com/gin-gonic/gin"
)

type OrderServer struct {
}

func (o *OrderServer) RegisterRouter(e *gin.Engine) {
	e.POST("/order", api.Handler()(OrderAPI))
}
