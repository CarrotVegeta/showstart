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
	"time"
)

func GetActivityDetail(activityID string, stFlpv string) map[string]interface{} {
	g := &models.GetActivityDetails{}
	g.Action = pkg.GetActivityDetailsAction
	g.Method = "POST"
	g.Query.ActivityId = activityID
	g.Query.Coupon = ""
	g.Query.StFlpv = stFlpv
	g.Query.TrackPath = ""
	g.Query.Terminal = pkg.Terminal
	bs, _ := jsoniter.Marshal(g.Query)
	g.Query.Sign = utils.GeneMD5(string(bs))
	g.Qtime = time.Now().UnixNano() / 1e6
	g.Ranstr = utils.RandStr(8)
	bs, _ = jsoniter.Marshal(g)
	c, err := utils.AesECBEncrypt(bs, []byte(pkg.Secret))
	if err != nil {
		log.Fatal(err.Error())
	}
	d := base64.StdEncoding.EncodeToString(c)
	bp := GetBodyParam(d)
	m, err := request.HttpDo(pkg.ActivityDetail, bp)
	if err != nil {
		log.Fatal(err.Error())
	}
	return m
}

type OrderConfig struct {
	GoodsId           int    //activityID
	SkuId             string //ticketID
	Price             int    //价格
	CommonPerformerID int    //观演人ID
	TelePhone         string //联系电话
	SessionId         int
	StFlpv            string
}

func Order(oc *OrderConfig) map[string]interface{} {
	geo := &models.GenerateOrder{}
	geo.Action = pkg.OrderAction
	geo.Method = "POST"
	od := &models.OrderDetails{
		GoodsType: 1,
		SkuType:   2,
		Num:       "1",
		GoodsId:   oc.GoodsId,
		SkuId:     oc.SkuId,
		Price:     oc.Price,
	}
	geo.Query.OrderDetails = append(geo.Query.OrderDetails, od)
	geo.Query.CommonPerfomerIds = append(geo.Query.CommonPerfomerIds, oc.CommonPerformerID)
	geo.Query.AreaCode = "86_CN"
	geo.Query.Telephone = oc.TelePhone
	geo.Query.CustomerName = "李川"
	geo.Query.ProvinceName = "四川"
	geo.Query.CityName = "成都"
	geo.Query.Address = "双流区梓州大道恒大名都"
	geo.Query.SessionId = oc.SessionId
	geo.Query.AmountPayable = fmt.Sprintf("%d.00", oc.Price)
	geo.Query.TotalAmount = oc.Price
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
	bp := GetBodyParam(d)
	e, err := request.HttpDo(pkg.Order2, bp)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
func GetBodyParam(d string) *models.BodyParam {
	bp := &models.BodyParam{
		Data:     d,
		Appid:    pkg.APPID,
		Terminal: pkg.Terminal,
		Version:  pkg.Version,
	}
	bs, _ := jsoniter.Marshal(bp)
	bp.Sign = utils.GeneMD5(string(bs))
	return bp
}
