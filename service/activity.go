package service

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
)

func GetActivityDetail(activityID string) map[string]interface{} {
	g := &models.GetActivityDetails{}
	g.ActivityId = activityID
	g.Coupon = ""
	g.StFlpv = config.Conf.StFlpv
	g.TrackPath = ""
	g.Terminal = pkg.Terminal
	return RequestWithBodyParam(pkg.GetActivityDetailsAction, pkg.A, "POST", g)
}
