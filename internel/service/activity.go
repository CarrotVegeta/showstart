package service

import (
	"github.com/CarrotVegeta/showstart/api"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetActivityListAPI(c *gin.Context, reply *api.Reply) {
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
	Search(activity.KeyWord, activity.PagNo, reply)
	return
}
func GetActivityDetailAPI(c *gin.Context, reply *api.Reply) {
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
	GetActivityDetail(activity.ActivityID, reply)
	return
}

func GetActivityDetail(activityID int, reply *api.Reply) {
	g := &models.GetActivityDetails{}
	g.ActivityId = strconv.Itoa(activityID)
	g.Coupon = ""
	g.StFlpv = api.User.StFlpv
	g.TrackPath = ""
	g.Terminal = pkg.Terminal
	result, err := RequestWithBodyParam(pkg.GetActivityDetailsAction, pkg.HW, http.MethodGet, g)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = result
}

func Search(keyword string, pageNo int, reply *api.Reply) {
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
		StFlpv:      api.User.StFlpv,
		Sign:        "",
		TrackPath:   "",
		Terminal:    pkg.Terminal,
	}
	result, err := RequestWithBodyParam(pkg.GetActivityListAction, pkg.HW, http.MethodGet, s)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	reply.Data = result
}
