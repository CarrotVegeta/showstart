package router

import (
	. "github.com/CarrotVegeta/showstart/controller"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/gin-gonic/gin"
)

type Ticket struct {
}

func (a *Ticket) RegisterRouter(e *gin.Engine) {
	e.POST("/tickets", server.Handler()(GetTicketList))
	e.POST("/ticket", server.Handler()(GetTicket))
}
