package service

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
)

func Search(keyword string, pageNo int) map[string]interface{} {
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
		StFlpv:      config.Conf.StFlpv,
		Sign:        "",
		TrackPath:   "",
		Terminal:    pkg.Terminal,
	}
	return RequestWithBodyParam(pkg.GetActivityListAction, pkg.B, "GET", s)
}
