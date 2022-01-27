package service

import (
	"fmt"
	"testing"
)

func TestGetActivityDetail(t *testing.T) {
	res := GetActivityDetail("170309")
	result := res["result"].(map[string]interface{})
	activityID := result["activityId"]
	fmt.Println(activityID)
}
