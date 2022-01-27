package engine

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/service"
	jsoniter "github.com/json-iterator/go"
)

type Spectator struct {
	ID     int
	Name   string
	UserID int
}

func getSpectatorList(userId int) *Spectator {
	res := service.GetSpectatorList()
	result := res["result"].([]interface{})
	for _, r := range result {
		bs, _ := jsoniter.Marshal(r)
		var s Spectator
		jsoniter.Unmarshal(bs, &s)
		user := fmt.Sprintf("id:%d 姓名:%s userId:%d", s.ID, s.Name, s.UserID)
		fmt.Println(user)
		if s.UserID == userId {
			return &s
		}
	}
	return nil
}
