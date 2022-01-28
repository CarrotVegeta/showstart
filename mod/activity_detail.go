package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
)

type ActivityDetail struct {
	ActivityID string
}

const activityDetailMod = "getActivityDetail"

func (ad *ActivityDetail) Open(s *engine.Server) {
	s.RegisterMod(ad, activityDetailMod)
	ad.ActivityID = fmt.Sprintf("%d", config.Conf.ActivityId)
}
func (ad *ActivityDetail) Request() {
	res := service.GetActivityDetail(ad.ActivityID)
	result := res["result"].(map[string]interface{})
	activityID := result["activityId"]
	fmt.Println(activityID)
}
