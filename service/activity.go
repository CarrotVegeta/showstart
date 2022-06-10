package service

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/server"
)

func GetActivityDetail(activityID string, reply *server.Reply) {
	g := &models.GetActivityDetails{}
	g.ActivityId = activityID
	g.Coupon = ""
	g.StFlpv = config.Conf.ShowStart.StFlpv
	g.TrackPath = ""
	g.Terminal = pkg.Terminal
	result, err := RequestWithBodyParam(pkg.GetActivityDetailsAction, pkg.A, "POST", g)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = result
}
