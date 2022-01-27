package engine

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/service"
)

type Ticket struct {
	SessionId  int
	BuyType    int
	TicketId   string
	CostPrice  string
	GoodType   int
	TicketNum  int
	ActivityId int
}

func getTicketList(activityId string, ticketId string) *Ticket {
	m := service.GetTicketList(activityId)
	if m != nil {
		result := m["result"].([]interface{})
		for _, r := range result {
			v := r.(map[string]interface{})
			ssId := fmt.Sprintf("sessionId:%d", (int)(v["sessionId"].(float64)))
			fmt.Println(ssId)
			sessionName := fmt.Sprintf("sessionName:%s", v["sessionName"].(string))
			fmt.Println(sessionName)
			ticketList := v["ticketList"].([]interface{})
			for _, t := range ticketList {
				tm := t.(map[string]interface{})
				msg := fmt.Sprintf("ticketType:%s ticketNum:%d ticketId:%s costPrice:%s showTime:%s",
					tm["ticketType"].(string),
					(int)(tm["ticketNum"].(float64)),
					tm["ticketId"].(string),
					tm["costPrice"].(string),
					tm["showTime"].(string))
				fmt.Println(msg)
				if tm["ticketId"].(string) == ticketId {
					return &Ticket{
						SessionId:  (int)(v["sessionId"].(float64)),
						BuyType:    (int)(tm["buyType"].(float64)),
						TicketId:   tm["ticketId"].(string),
						CostPrice:  tm["costPrice"].(string),
						GoodType:   (int)(tm["goodType"].(float64)),
						TicketNum:  (int)(tm["ticketNum"].(float64)),
						ActivityId: (int)(tm["activityId"].(float64)),
					}
				}
			}
		}
	}
	return nil
}
