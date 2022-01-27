package service

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
)

func GetTicketList(activityId string) map[string]interface{} {
	t := &models.TicketParam{
		ActivityId: activityId,
		Coupon:     "",
		StFlpv:     config.Conf.StFlpv,
		TrackPath:  "",
		Terminal:   pkg.Terminal,
	}
	return RequestWithBodyParam(pkg.GetTicketListAction, pkg.Detail, "POST", t)
}
