package service

//
//type Spectator struct {
//	ID     int
//	Name   string
//	UserID int
//}
//
//const spectatorMod = "getSpectatorList"
//
//func (s *Spectator) Open(ss *engine.Server) {
//	ss.RegisterMod(s, spectatorMod)
//}
//func (s *Spectator) Request() {
//	s.GetSpectator()
//}
//func (s *Spectator) GetSpectator() *Spectator {
//	res := service.GetSpectatorList()
//	result := res["result"].([]interface{})
//	for _, r := range result {
//		bs, _ := jsoniter.Marshal(r)
//		var sp Spectator
//		err := jsoniter.Unmarshal(bs, &sp)
//		if err != nil {
//			log.Fatal(err.Error())
//			return nil
//		}
//		user := fmt.Sprintf("id:%d 姓名:%s userId:%d", s.ID, s.Name, s.UserID)
//		fmt.Println(user)
//		if s.UserID == sp.UserID {
//			return &sp
//		}
//	}
//	return nil
//}
//func GetSpectatorList() map[string]interface{} {
//	s := &models.RequestSpectator{
//		StFlpv:    conf.Conf.StFlpv,
//		Sign:      "",
//		TrackPath: "",
//		Terminal:  pkg.Terminal,
//	}
//	return RequestWithBodyParam(pkg.GetSpectatorAction, pkg.B, "POST", s)
//}
