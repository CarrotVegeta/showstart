package engine

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/service"
)

func order(activityId string, ticketId string) {
	t := getTicketList(activityId, ticketId)
	//getSpectatorList(1225930)
	oc := &service.OrderConfig{
		GoodsType:         t.GoodType,
		SkuType:           t.BuyType,
		Num:               fmt.Sprintf("%d", t.TicketNum),
		GoodsId:           t.ActivityId,
		SkuId:             t.TicketId,
		Price:             t.CostPrice,
		CommonPerformerID: 1225930,
		TelePhone:         "13118115155",
		SessionId:         t.SessionId,
		StFlpv:            config.Conf.StFlpv,
	}
	pa := &service.PersonAddress{
		Telephone:    oc.TelePhone,
		CustomerName: "李川",
		ProvinceName: "四川省",
		CityName:     "成都市",
		Address:      "梓州大道",
	}
	//ticker := time.NewTicker(time.Second)
	//for {
	//	select {
	//	case <-ticker.C:
	res := service.Order(oc, pa)
	fmt.Println(res)
	//	}
	//}
}
