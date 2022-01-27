package service

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"testing"
)

func TestGetActivityDetail(t *testing.T) {
	config.Init()
	res := GetActivityDetail("170309")
	result := res["result"].(map[string]interface{})
	activityID := result["activityId"]
	fmt.Println(activityID)
}
