package service

import (
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/server"
	"strconv"
)

func GetActivityDetail(activityID int, reply *server.Reply) {
	g := &models.GetActivityDetails{}
	g.ActivityId = strconv.Itoa(activityID)
	g.Coupon = ""
	g.StFlpv = server.User.StFlpv
	g.TrackPath = ""
	g.Terminal = pkg.Terminal
	result, err := RequestWithBodyParam(pkg.GetActivityDetailsAction, pkg.HW, "POST", g)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = result
}

func Search(keyword string, pageNo int, reply *server.Reply) {
	s := &models.RequestSearch{
		PageNo:      pageNo,
		CityCode:    "28",
		Keyword:     keyword,
		Style:       "",
		ActivityIds: "",
		CouponCode:  "",
		PerformerId: "",
		HosterId:    "",
		SiteId:      "",
		Tag:         "",
		TourId:      "",
		ThemeId:     "",
		StFlpv:      server.User.StFlpv,
		Sign:        "",
		TrackPath:   "",
		Terminal:    pkg.Terminal,
	}
	result, err := RequestWithBodyParam(pkg.GetActivityListAction, pkg.HW, "GET", s)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = result
}
