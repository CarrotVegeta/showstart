package controller

//
//type Order struct {
//	ActivityID        int
//	TicketID          string
//	CommonPerformerID int
//}
//
//const orderMod = "order"
//
//var num = 1
//var ch = make(chan bool)
//
//func (o *Order) Open(s *engine.Server) {
//	s.RegisterMod(o, orderMod)
//	o.ActivityID = config.Conf.ActivityId
//	o.TicketID = config.Conf.TicketId
//}
//func (o *Order) Request() {
//	t := &Ticket{}
//	t.ActivityId = o.ActivityID
//	t.TicketId = o.TicketID
//	t = t.GetTicketList()
//	if config.Conf.TicketNum > 0 {
//		num = config.Conf.TicketNum
//	}
//	oc := &service.OrderConfig{
//		GoodsType:         t.GoodType,
//		SkuType:           t.Type,
//		Num:               fmt.Sprintf("%d", num),
//		GoodsId:           t.ActivityId,
//		SkuId:             t.TicketId,
//		Price:             t.SellingPrice,
//		CommonPerformerID: o.CommonPerformerID,
//		TelePhone:         config.Conf.Telephone,
//		SessionId:         t.SessionId,
//		StFlpv:            config.Conf.StFlpv,
//	}
//	pa := &service.PersonAddress{
//		Telephone:    oc.TelePhone,
//		CustomerName: config.Conf.CustomerName,
//		ProvinceName: config.Conf.Province,
//		CityName:     config.Conf.CityName,
//		Address:      config.Conf.Address,
//	}
//	if config.Conf.CronTime == "" {
//		o.OrderNow(oc, pa)
//	} else {
//		o.Order(oc, pa)
//	}
//
//	var j int
//	for {
//		select {
//		case <-ch:
//			j++
//		}
//		if j == config.Conf.OrderNum {
//			break
//		}
//	}
//}
//func (o *Order) Order(oc *service.OrderConfig, pa *service.PersonAddress) {
//	c := cron.New()
//	err := c.AddFunc(config.Conf.CronTime, func() {
//		log.Println("开始抢票========================》")
//		var flag bool
//		for {
//			if flag {
//				i := rand.Intn(int(config.Conf.Ticker))
//				time.Sleep(time.Millisecond * (time.Duration(i)))
//			} else {
//				time.Sleep(time.Millisecond * 100)
//			}
//			flag = true
//			o.GoOrder(oc, pa)
//		}
//	})
//	if err != nil {
//		log.Fatal(err.Error())
//		return
//	}
//	c.Start()
//}
//func (o *Order) OrderNow(oc *service.OrderConfig, pa *service.PersonAddress) {
//	log.Println("开始抢票========================》")
//	var flag bool
//	for {
//		if flag {
//			i := rand.Intn(int(config.Conf.Ticker))
//			time.Sleep(time.Millisecond * (time.Duration(i)))
//		} else {
//			time.Sleep(time.Millisecond * 100)
//		}
//		flag = true
//		o.GoOrder(oc, pa)
//	}
//}
//func (o *Order) GoOrder(oc *service.OrderConfig, pa *service.PersonAddress) {
//	res := service.Order(oc, pa)
//	log.Println(res)
//	if res["success"].(bool) == true {
//		fmt.Println("===========下单中》》》》》》》》》》》》》》》》》》》》")
//		result := res["result"].(map[string]interface{})
//		orderJobKey := result["orderJobKey"].(string)
//		res1 := service.GetCoreOrderResult(orderJobKey, oc.StFlpv)
//		log.Println(res1)
//		if res1["success"].(bool) == true {
//			if _, ok := res1["result"]; ok {
//				if reflect.TypeOf(res1["result"]).Kind() == reflect.Map {
//					result1 := res1["result"].(map[string]interface{})
//					if result1["orderId"].(string) != "" {
//						fmt.Println("===================抢票成功=================================")
//						ch <- true
//					}
//				}
//			}
//		}
//	}
//}
