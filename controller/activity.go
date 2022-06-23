package controller

import (
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/service"
	"github.com/gin-gonic/gin"
)

func GetActivityList(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	activity := models.NewActivity()
	err := c.ShouldBindJSON(activity)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if activity.KeyWord == "" {
		reply.Error = "key_word is null"
		return
	}
	if activity.PagNo == 0 {
		activity.PagNo = 10
	}
	service.Search(activity.KeyWord, activity.PagNo, reply)
	return
}
func GetActivityDetail(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	activity := models.NewActivity()
	err := c.ShouldBindJSON(activity)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if activity.ActivityID == 0 {
		reply.Error = "activity_id is 0"
		return
	}
	service.GetActivityDetail(activity.ActivityID, reply)
	return
}
