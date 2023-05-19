package router

import (
	"github.com/CarrotVegeta/showstart/api"
	. "github.com/CarrotVegeta/showstart/pkg/service"
	"github.com/gin-gonic/gin"
)

type Ticket struct {
}

func (a *Ticket) RegisterRouter(e *gin.Engine) {
	e.POST("/tickets", api.Handler(GetTicketListAPI))
	e.POST("/ticket", api.Handler(GetTicketAPI))
}
