package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type Spectator struct {
	ID     int
	Name   string
	UserID int
}

const spectatorMod = "getSpectatorList"

func (s *Spectator) Open(ss *engine.Server) {
	ss.RegisterMod(s, spectatorMod)
}
func (s *Spectator) Request() {
	s.GetSpectator()
}
func (s *Spectator) GetSpectator() *Spectator {
	res := service.GetSpectatorList()
	result := res["result"].([]interface{})
	for _, r := range result {
		bs, _ := jsoniter.Marshal(r)
		var sp Spectator
		err := jsoniter.Unmarshal(bs, &sp)
		if err != nil {
			log.Fatal(err.Error())
			return nil
		}
		user := fmt.Sprintf("id:%d 姓名:%s userId:%d", s.ID, s.Name, s.UserID)
		fmt.Println(user)
		if s.UserID == sp.UserID {
			return &sp
		}
	}
	return nil
}
