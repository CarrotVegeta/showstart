package controller

import (
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/service"
	"github.com/gin-gonic/gin"
)

func GetTicketList(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	ticket := models.NewTicket()
	if err := c.ShouldBindJSON(ticket); err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if ticket.ActivityId == 0 {
		reply.Error = "activity_id is 0"
		return
	}
	tickets, err := service.GetTicketList(ticket.ActivityId, "")
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = tickets
	return
}
func GetTicket(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	ticket := models.NewTicket()
	if err := c.ShouldBindJSON(ticket); err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if ticket.ActivityId == 0 {
		reply.Error = "activity_id is 0"
		return
	}
	if ticket.TicketId == "" {
		reply.Error = "ticket_id is null"
		return
	}
	tickets, err := service.GetTicketList(ticket.ActivityId, ticket.TicketId)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = tickets
	return
}
