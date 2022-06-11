package router

import (
	. "github.com/CarrotVegeta/showstart/controller"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/gin-gonic/gin"
)

type Activity struct {
}

func (a *Activity) RegisterRouter(e *gin.Engine) {
	e.POST("/activity/detail", server.Handler()(GetActivityDetail))
	e.POST("/activity/list", server.Handler()(GetActivityList))
}
