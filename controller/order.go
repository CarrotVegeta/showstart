package controller

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/service"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
	"math/rand"
	"reflect"
	"time"
)

var (
	num    = 1
	ticker = 500
)
var ch = make(chan bool)
var cherr = make(chan error)

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
	tickets, err := service.GetTicketList(order.ActivityID, order.TicketID)
	if err != nil {
		reply.Error = err.Error()
		logger.FileLog.Error(err.Error())
		return
	}
	if order.TicketNum > 0 {
		num = order.TicketNum
	}
	if order.Ticker == 0 {
		order.Ticker = int64(ticker)
	}
	t := tickets[0]
	oc := &service.OrderConfig{
		GoodsType:         t.GoodType,
		SkuType:           t.Type,
		Num:               fmt.Sprintf("%d", num),
		GoodsId:           t.ActivityId,
		SkuId:             t.TicketId,
		Price:             t.SellingPrice,
		CommonPerformerID: order.CommonPerformerID,
		TelePhone:         order.Telephone,
		SessionId:         t.SessionId,
		StFlpv:            server.User.StFlpv,
	}
	pa := &service.PersonAddress{
		Telephone:    oc.TelePhone,
		CustomerName: order.CustomerName,
		ProvinceName: order.ProvinceName,
		CityName:     order.CityName,
		Address:      order.Address,
	}
	if order.CronTime == "" {
		orderNow(oc, pa, order)
	} else {
		err := cronOrder(oc, pa, order)
		if err != nil {
			reply.Error = err.Error()
			logger.FileLog.Error(err.Error())
			return
		}
	}
	var j int
	for {
		select {
		case <-ch:
			j++
		case err := <-cherr:
			reply.Error = err.Error()
			logger.FileLog.Error(err.Error())
			close(cherr)
			close(ch)
		}
		if j == order.OrderNum {
			reply.Data = "抢票成功"
			break
		}
	}
	return
}

func cronOrder(oc *service.OrderConfig, pa *service.PersonAddress, o models.OrderModel) error {
	c := cron.New()
	err := c.AddFunc(o.CronTime, func() {
		log.Println("开始抢票========================》")
		var flag bool
		for {
			if flag {
				i := rand.Intn(int(o.Ticker))
				time.Sleep(time.Millisecond * (time.Duration(i)))
			} else {
				time.Sleep(time.Millisecond * 100)
			}
			flag = true
			goOrder(oc, pa)
		}
	})
	if err != nil {
		return fmt.Errorf("start cron err :%v", err.Error())
	}
	c.Start()
	return nil
}
func orderNow(oc *service.OrderConfig, pa *service.PersonAddress, o models.OrderModel) {
	log.Println("开始抢票========================》")
	var flag bool
	for {
		if flag {
			i := rand.Intn(int(o.Ticker))
			time.Sleep(time.Millisecond * (time.Duration(i)))
		} else {
			time.Sleep(time.Millisecond * 100)
		}
		flag = true
		goOrder(oc, pa)
	}
}
func goOrder(oc *service.OrderConfig, pa *service.PersonAddress) {
	res, err := service.Order(oc, pa)
	if err != nil {
		cherr <- err
	}
	if res["success"].(bool) == true {
		result := res["result"].(map[string]interface{})
		orderJobKey := result["orderJobKey"].(string)
		res1, err := service.GetCoreOrderResult(orderJobKey)
		if err != nil {
			cherr <- err
		}
		if res1["success"].(bool) == true {
			if _, ok := res1["result"]; ok {
				if reflect.TypeOf(res1["result"]).Kind() == reflect.Map {
					result1 := res1["result"].(map[string]interface{})
					if result1["orderId"].(string) != "" {
						fmt.Println("===================抢票成功=================================")
						ch <- true
					}
				}
			}
		}
	}
}
