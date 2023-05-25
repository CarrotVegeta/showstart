package service

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/api"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

var (
	num = 1
)

func OrderAPI(c *gin.Context) (reply *api.Reply) {
	reply = api.NewReply()
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
	oc := &OrderConfig{
		GoodsType:         order.GoodsType,
		SkuType:           order.SkuType,
		Num:               fmt.Sprintf("%d", num),
		GoodsId:           order.ActivityID,
		SkuId:             order.TicketID,
		Price:             order.SellingPrice,
		CommonPerformerID: order.CommonPerformerID,
		TelePhone:         order.Telephone,
		SessionId:         order.SessionId,
		StFlpv:            api.User.StFlpv,
	}
	pa := &PersonAddress{
		Telephone:    oc.TelePhone,
		CustomerName: order.CustomerName,
		ProvinceName: order.ProvinceName,
		CityName:     order.CityName,
		Address:      order.Address,
	}
	tickets, err := GetTicketList(order.ActivityID, order.TicketID)
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

func goOrder(oc *OrderConfig, pa *PersonAddress) (bool, string, error) {
	resOrder, err := Order(oc, pa)
	if err != nil {
		return false, "", fmt.Errorf("goorder error :%v", err.Error())
	}
	if _, ok := resOrder["success"]; !ok || !resOrder["success"].(bool) {
		return false, resOrder["msg"].(string), nil
	}
	result := resOrder["result"].(map[string]interface{})
	orderJobKey := result["orderJobKey"].(string)
	resp, err := GetCoreOrderResult(orderJobKey)
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

type OrderConfig struct {
	GoodsType         int
	SkuType           int
	Num               string
	GoodsId           int    //activityID
	SkuId             string //ticketID
	Price             string //价格
	CommonPerformerID int    //观演人ID
	TelePhone         string //联系电话
	SessionId         int
	StFlpv            string
}
type PersonAddress struct {
	Telephone    string //电话
	CustomerName string //姓名
	ProvinceName string //省份
	CityName     string //城市
	Address      string //详细地址

}

func GetCoreOrderResult(orderJobKey string) (map[string]interface{}, error) {
	gc := &models.CoreOrderResult{
		OrderJobKey: orderJobKey,
		StFlpv:      api.User.StFlpv,
		TrackPath:   "",
		Terminal:    pkg.Terminal,
	}
	return RequestWithBodyParam(pkg.GetCoreOrderResultAction, pkg.APPNJ, http.MethodPost, gc)
}
func Order(oc *OrderConfig, pa *PersonAddress) (map[string]interface{}, error) {
	geo := &models.GenerateOrder{}
	price, _ := strconv.ParseFloat(oc.Price, 64)
	od := &models.OrderDetails{
		GoodsType: oc.GoodsType,
		SkuType:   oc.SkuType,
		Num:       oc.Num,
		GoodsId:   oc.GoodsId,
		SkuId:     oc.SkuId,
		Price:     price,
	}
	geo.OrderDetails = append(geo.OrderDetails, od)
	geo.CommonPerfomerIds = append(geo.CommonPerfomerIds, oc.CommonPerformerID)
	geo.AreaCode = "86_CN"
	geo.Telephone = pa.Telephone
	geo.CustomerName = pa.CustomerName
	geo.ProvinceName = pa.ProvinceName
	geo.CityName = pa.CityName
	geo.Address = pa.Address
	geo.SessionId = oc.SessionId
	geo.AmountPayable = fmt.Sprintf("%.2f", price)
	geo.TotalAmount = price
	geo.StFlpv = oc.StFlpv
	geo.Terminal = pkg.Terminal
	return RequestWithBodyParam(pkg.OrderAction, pkg.APPNJ, http.MethodPost, geo)
}
