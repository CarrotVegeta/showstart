package mod

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/service"
	"log"
)

type Ticket struct {
	SessionId    int
	Type         int
	TicketId     string
	SellingPrice string
	GoodType     int
	TicketNum    int
	ActivityId   int
	MemberNum    int
	RemainTicket int
}

func (t *Ticket) Request() {
	t.GetTicketList()
}

const ticketListMod = "getTicketList"

func (t *Ticket) Open(s *engine.Server) {
	s.RegisterMod(t, ticketListMod)
	t.ActivityId = config.Conf.ActivityId
}
func (t *Ticket) GetTicketList() *Ticket {
	m := service.GetTicketList(fmt.Sprintf("%d", t.ActivityId))
	if m != nil {
		result := m["result"].([]interface{})
		for _, r := range result {
			v := r.(map[string]interface{})
			ssId := fmt.Sprintf("sessionId:%d", (int)(v["sessionId"].(float64)))
			log.Println(ssId)
			sessionName := fmt.Sprintf("sessionName:%s", v["sessionName"].(string))
			log.Println(sessionName)
			ticketList := v["ticketList"].([]interface{})
			for _, tl := range ticketList {
				tm := tl.(map[string]interface{})
				msg := fmt.Sprintf("ticketType:%s ticketNum:%d ticketId:%s sellingPrice:%s showTime:%s",
					tm["ticketType"].(string),
					(int)(tm["ticketNum"].(float64)),
					tm["ticketId"].(string),
					tm["sellingPrice"].(string),
					tm["showTime"].(string))
				log.Println(msg)
				if tm["ticketId"].(string) == t.TicketId {
					return &Ticket{
						SessionId:    (int)(v["sessionId"].(float64)),
						Type:         (int)(tm["type"].(float64)),
						TicketId:     tm["ticketId"].(string),
						SellingPrice: tm["sellingPrice"].(string),
						GoodType:     (int)(tm["goodType"].(float64)),
						TicketNum:    (int)(tm["ticketNum"].(float64)),
						ActivityId:   (int)(tm["activityId"].(float64)),
						MemberNum:    (int)(tm["memberNum"].(float64)),
					}
				}
			}
		}
	}
	return nil
}
