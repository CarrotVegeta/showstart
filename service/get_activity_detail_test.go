package service

import (
	"fmt"
	"testing"
)

func TestGetActivityDetail(t *testing.T) {
	res := GetActivityDetail("173034")
	result := res["result"].(map[string]interface{})
	activityID := result["activityId"]
	fmt.Println(activityID)
}
