package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
	"time"
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
	ticker := time.NewTicker(time.Millisecond * (time.Duration)(config.Conf.Ticker))
	for {
		select {
		case <-ticker.C:
			res := service.Order(oc, pa)
			fmt.Println(res)
			if res["success"].(bool) == true {
				fmt.Println("===========下单中》》》》》》》》》》》》》》》》》》》》")
				result := res["result"].(map[string]interface{})
				orderJobKey := result["orderJobKey"].(string)
				res1 := service.GetCoreOrderResult(orderJobKey, oc.StFlpv)
				if res1["success"].(bool) == true {
					if _, ok := res1["result"]; ok {
						result1 := res1["result"].(map[string]interface{})
						if result1["orderId"].(string) != "" {
							fmt.Println("===================抢票成功=================================")
							ticker.Stop()
							return
						}
					}
				}
			}
		}
	}
}
