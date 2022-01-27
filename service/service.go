package service

import (
	"encoding/base64"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/request"
	"github.com/CarrotVegeta/showstart/utils"
	jsoniter "github.com/json-iterator/go"
	"log"
	"time"
)

func RequestWithBodyParam(action, url string, method string, query interface{}) map[string]interface{} {
	p := &models.Param{
		Action: action,
		Method: method,
		Query:  query,
		Body:   nil,
		Qtime:  time.Now().UnixNano() / 1e6,
		Ranstr: utils.RandStr(8),
	}
	bs, _ := jsoniter.Marshal(p)
	c, err := utils.AesECBEncrypt(bs, []byte(pkg.Secret))
	if err != nil {
		log.Fatal(err.Error())
	}
	d := base64.StdEncoding.EncodeToString(c)
	bp := GetBodyParam(d, string(bs))
	rm, err := request.HttpDo(url, bp)
	if err != nil {
		log.Fatal(err.Error())
	}
	return rm
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

func GetBodyParam(d, s string) *models.BodyParam {
	bp := &models.BodyParam{
		Data:     d,
		Appid:    pkg.APPID,
		Terminal: pkg.Terminal,
		Version:  pkg.Version,
	}
	bp.Sign = utils.GeneMD5(s)
	return bp
}
