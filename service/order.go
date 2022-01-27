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
	bp := GetBodyParam(d, string(bs))
	e, err := request.HttpDo(pkg.Order2, bp)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
