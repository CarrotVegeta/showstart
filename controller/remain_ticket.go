package controller

//
//type TicketRemain struct {
//	TicketId   string
//	ActivityId int
//}
//
//func (tr *TicketRemain) Request() {
//	tr.GetTicketRemain()
//}
//
//const ticketRemainMod = "getRemainTicket"
//
//func (tr *TicketRemain) Open(s *engine.Server) {
//	s.RegisterMod(tr, ticketRemainMod)
//	tr.ActivityId = config.Conf.ActivityId
//	tr.TicketId = config.Conf.TicketId
//}
//
//var chb = make(chan int)
//
//func (tr *TicketRemain) GetTicketRemain() {
//	log.Println("开始查询余票===============================》")
//	go func() {
//		var flag bool
//		for {
//			if flag {
//				i := rand.Intn(int(config.Conf.Ticker))
//				time.Sleep(time.Millisecond * (time.Duration(i)))
//			}
//			go tr.getRemainTicket()
//			flag = true
//		}
//	}()
//	select {
//	case f := <-chb:
//		log.Printf("有余票了 剩余%d张", f)
//	}
//}
//func (tr *TicketRemain) getRemainTicket() {
//	m := service.GetTicketList(fmt.Sprintf("%d", tr.ActivityId))
//	if m != nil {
//		result := m["result"].([]interface{})
//		for _, r := range result {
//			v := r.(map[string]interface{})
//			ticketList := v["ticketList"].([]interface{})
//			for _, tl := range ticketList {
//				tm := tl.(map[string]interface{})
//				remainTicket := (int)(tm["remainTicket"].(float64))
//				log.Printf("ticketType:%s ticketNum:%d ticketId:%s sellingPrice:%s showTime:%s remainTicket:%d \n",
//					tm["ticketType"].(string),
//					(int)(tm["ticketNum"].(float64)),
//					tm["ticketId"].(string),
//					tm["sellingPrice"].(string),
//					tm["showTime"].(string),
//					remainTicket)
//				if tm["ticketId"].(string) == tr.TicketId && remainTicket > 0 {
//					chb <- remainTicket
//				}
//			}
//		}
//	}
//}
