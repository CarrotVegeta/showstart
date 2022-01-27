package engine

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/service"
)

func getActivityList(keyword string, pagno int) {
	res := service.Search(keyword, pagno)
	result := res["result"].(map[string]interface{})
	activityInfo := result["activityInfo"].([]interface{})
	for _, v := range activityInfo {
		m := v.(map[string]interface{})
		sprintf := fmt.Sprintf("activityId:%d   title: %s ", (int)(m["activityId"].(float64)), m["title"].(string))
		fmt.Println(sprintf)
	}
}

func getActivityDetail(activityId string) {
	res := service.GetActivityDetail(activityId)
	result := res["result"].(map[string]interface{})
	activityID := result["activityId"]
	fmt.Println(activityID)
}
