package controller

import (
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/service"
	"github.com/gin-gonic/gin"
)

func GetActivityDetail(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	activity := models.NewActivity()
	err := c.ShouldBindJSON(activity)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if activity.ActivityID == "" {
		reply.Error = "activity is null"
		return
	}
	service.GetActivityDetail(activity.ActivityID, reply)
	return
}
