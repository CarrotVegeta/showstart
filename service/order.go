package service

import (
	"encoding/base64"
	"fmt"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/request"
	"github.com/CarrotVegeta/showstart/utils"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strconv"
	"time"
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
	geo.Action = pkg.OrderAction
	geo.Method = "POST"
	price, _ := strconv.ParseFloat(oc.Price, 64)
	od := &models.OrderDetails{
		GoodsType: oc.GoodsType,
		SkuType:   oc.SkuType,
		Num:       oc.Num,
		GoodsId:   oc.GoodsId,
		SkuId:     oc.SkuId,
		Price:     (int)(price),
	}
	geo.Query.OrderDetails = append(geo.Query.OrderDetails, od)
	geo.Query.CommonPerfomerIds = append(geo.Query.CommonPerfomerIds, oc.CommonPerformerID)
	geo.Query.AreaCode = "86_CN"
	geo.Query.Telephone = pa.Telephone
	geo.Query.CustomerName = pa.CustomerName
	geo.Query.ProvinceName = pa.ProvinceName
	geo.Query.CityName = pa.CityName
	geo.Query.Address = pa.Address
	geo.Query.SessionId = oc.SessionId
	geo.Query.AmountPayable = fmt.Sprintf("%.2f", price)
	geo.Query.TotalAmount = (int)(price)
	geo.Query.StFlpv = oc.StFlpv
	geo.Query.Terminal = pkg.Terminal
	geo.Qtime = time.Now().UnixNano() / 1e6
	geo.Ranstr = utils.RandStr(8)
	bs, _ := jsoniter.Marshal(geo)
	c, err := utils.AesECBEncrypt(bs, []byte(pkg.Secret))
	if err != nil {
		log.Fatal(err.Error())
	}
	d := base64.StdEncoding.EncodeToString(c)
	bp := GetBodyParam(d, string(bs))
	e, err := request.HttpDo(pkg.GetOrderUrl(oc.SkuType), bp)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
