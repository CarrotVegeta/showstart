package service

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	res := Search("元气森林", 1)
	fmt.Println(res)
	result := res["result"].(map[string]interface{})
	activityInfo := result["activityInfo"].([]interface{})
	for _, v := range activityInfo {
		m := v.(map[string]interface{})
		sprintf := fmt.Sprintf("activityId:%d   title: %s ", (int)(m["activityId"].(float64)), m["title"].(string))
		fmt.Println(sprintf)
	}
}
