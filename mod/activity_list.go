package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
)

type ActivityList struct {
	Keyword string
	Pagno   int
}

const activityListMod = "getActivityList"

func (al *ActivityList) Open(s *engine.Server) {
	s.RegisterMod(al, activityListMod)
	al.Keyword = config.Conf.Keyword
	al.Pagno = config.Conf.Pagno
}

func (al *ActivityList) Request() {
	res := service.Search(al.Keyword, al.Pagno)
	result := res["result"].(map[string]interface{})
	activityInfo := result["activityInfo"].([]interface{})
	for _, v := range activityInfo {
		m := v.(map[string]interface{})
		sprintf := fmt.Sprintf("activityId:%d   title: %s ", (int)(m["activityId"].(float64)), m["title"].(string))
		fmt.Println(sprintf)
	}
}
