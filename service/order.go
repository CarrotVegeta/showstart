package service

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"strconv"
)

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

func Order(oc *OrderConfig, pa *PersonAddress) map[string]interface{} {
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
	return RequestWithBodyParam(pkg.OrderAction, pkg.GetOrderUrl(oc.SkuType), "POST", geo)
}
