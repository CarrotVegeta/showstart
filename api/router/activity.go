package router

import (
	"github.com/CarrotVegeta/showstart/api"
	. "github.com/CarrotVegeta/showstart/pkg/service"
	"github.com/gin-gonic/gin"
)

type Activity struct {
}

func (a *Activity) RegisterRouter(e *gin.Engine) {
	e.POST("/activity/detail", api.Handler()(GetActivityDetailAPI))
	e.POST("/activity/list", api.Handler()(GetActivityListAPI))
}
