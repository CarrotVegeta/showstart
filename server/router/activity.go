package router

import "github.com/gin-gonic/gin"

type Activity struct {
}

func (a *Activity) RegisterRouter(e *gin.Engine) {
	e.POST("/")
}
