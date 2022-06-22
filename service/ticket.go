package service

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/server"
	"strconv"
)

func GetTicketList(activityId int, ticketID string) ([]models.TicketModel, error) {
	activityIdStr := strconv.FormatInt(int64(activityId), 10)
	t := &models.TicketParam{
		ActivityId: activityIdStr,
		Coupon:     "",
		StFlpv:     server.User.StFlpv,
		TrackPath:  "",
		Terminal:   pkg.Terminal,
	}
	m, err := RequestWithBodyParam(pkg.GetTicketListAction, pkg.A, "POST", t)
	if err != nil {
		return nil, fmt.Errorf("getticketlist request err:%v", err.Error())
	}
	var tickets []models.TicketModel
	if m != nil {
		if _, ok := m["result"]; !ok {
			return nil, fmt.Errorf("get ticket_list request error:%v", m["msg"].(string))
		}
		result := m["result"].([]interface{})
		for _, r := range result {
			v := r.(map[string]interface{})
			ticketList := v["ticketList"].([]interface{})
			for _, tl := range ticketList {
				tm := tl.(map[string]interface{})
				ticket := models.NewTicket()
				ticket.SessionId = (int)(v["sessionId"].(float64))
				ticket.Type = (int)(tm["type"].(float64))
				ticket.SellingPrice = tm["sellingPrice"].(string)
				ticket.GoodType = (int)(tm["goodType"].(float64))
				ticket.TicketNum = (int)(tm["ticketNum"].(float64))
				ticket.MemberNum = (int)(tm["memberNum"].(float64))
				ticket.TicketType = tm["ticketType"].(string)
				ticket.ShowTime = tm["showTime"].(string)
				ticket.TicketId = tm["ticketId"].(string)
				if ticketID != "" && ticketID == ticket.TicketId {
					tickets = append(tickets, ticket)
					break
				}
				if ticketID == "" {
					tickets = append(tickets, ticket)
				}
			}
		}
	}
	return tickets, nil
}
