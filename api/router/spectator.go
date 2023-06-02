package router

import (
	"github.com/CarrotVegeta/showstart/api"
	. "github.com/CarrotVegeta/showstart/internel/service"
	"github.com/gin-gonic/gin"
)

type SpectatorServer struct {
}

func (o *SpectatorServer) RegisterRouter(e *gin.Engine) {
	e.POST("/spectators", api.Handler(GetSpectators))
}
