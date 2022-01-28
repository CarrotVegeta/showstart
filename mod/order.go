package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
)

type Order struct {
	ActivityID        int
	TicketID          string
	CommonPerformerID int
}

const orderMod = "order"

func (o *Order) Open(s *engine.Server) {
	s.RegisterMod(o, orderMod)
	o.ActivityID = config.Conf.ActivityId
	o.TicketID = config.Conf.TicketId
}
func (o *Order) Request() {
	t := &Ticket{}
	t.ActivityId = o.ActivityID
	t.TicketId = o.TicketID
	t = t.GetTicketList()
	oc := &service.OrderConfig{
		GoodsType:         t.GoodType,
		SkuType:           t.Type,
		Num:               fmt.Sprintf("%d", t.MemberNum),
		GoodsId:           t.ActivityId,
		SkuId:             t.TicketId,
		Price:             t.SellingPrice,
		CommonPerformerID: o.CommonPerformerID,
		TelePhone:         config.Conf.Telephone,
		SessionId:         t.SessionId,
		StFlpv:            config.Conf.StFlpv,
	}
	pa := &service.PersonAddress{
		Telephone:    oc.TelePhone,
		CustomerName: config.Conf.CustomerName,
		ProvinceName: config.Conf.Province,
		CityName:     config.Conf.CityName,
		Address:      config.Conf.Address,
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
