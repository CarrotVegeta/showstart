package controller

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/service"
	"github.com/gin-gonic/gin"
	"reflect"
)

var (
	num = 1
)

func Order(c *gin.Context) (reply *server.Reply) {
	reply = server.NewReply()
	order := models.NewOrder()
	if err := c.ShouldBindJSON(order); err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if order.ActivityID == 0 {
		reply.Error = "activity_id is 0"
		return
	}
	if order.TicketID == "" {
		reply.Error = "ticket_id is null"
		return
	}
	if order.TicketNum > 0 {
		num = order.TicketNum
	}
	oc := &service.OrderConfig{
		GoodsType:         order.GoodsType,
		SkuType:           order.SkuType,
		Num:               fmt.Sprintf("%d", num),
		GoodsId:           order.ActivityID,
		SkuId:             order.TicketID,
		Price:             order.SellingPrice,
		CommonPerformerID: order.CommonPerformerID,
		TelePhone:         order.Telephone,
		SessionId:         order.SessionId,
		StFlpv:            server.User.StFlpv,
	}
	pa := &service.PersonAddress{
		Telephone:    oc.TelePhone,
		CustomerName: order.CustomerName,
		ProvinceName: order.ProvinceName,
		CityName:     order.CityName,
		Address:      order.Address,
	}
	tickets, err := service.GetTicketList(order.ActivityID, order.TicketID)
	if err != nil {
		reply.Error = err.Error()
		return
	}
	ticket := tickets[0]
	if ticket.RemainTicket <= 0 {
		reply.Data = "库存不足,请稍后再试!"
		return
	}
	isSuccess, msg, err := goOrder(oc, pa)
	if err != nil {
		reply.Error = err.Error()
		return
	}
	if !isSuccess {
		reply.Data = msg
		return
	}
	reply.Data = "抢票成功"
	return
}

func goOrder(oc *service.OrderConfig, pa *service.PersonAddress) (bool, string, error) {
	resOrder, err := service.Order(oc, pa)
	if err != nil {
		return false, "", fmt.Errorf("goorder error :%v", err.Error())
	}
	if _, ok := resOrder["success"]; !ok || !resOrder["success"].(bool) {
		return false, resOrder["msg"].(string), nil
	}
	result := resOrder["result"].(map[string]interface{})
	orderJobKey := result["orderJobKey"].(string)
	resp, err := service.GetCoreOrderResult(orderJobKey)
	if err != nil {
		return false, "", fmt.Errorf("get core order result error:%v", err.Error())
	}
	if _, ok := resp["success"]; !ok || !resp["success"].(bool) {
		return false, resp["msg"].(string), nil
	}
	if _, ok := resp["result"]; ok {
		if reflect.TypeOf(resp["result"]).Kind() == reflect.Map {
			result1 := resp["result"].(map[string]interface{})
			if result1["orderId"].(string) != "" {
				logger.FileLog.Info("抢票成功")
				return true, "", nil
			}
		}
	}
	return true, "", nil
}
