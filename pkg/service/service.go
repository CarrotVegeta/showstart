package service

import (
	"encoding/base64"
	"fmt"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/CarrotVegeta/showstart/pkg/request"
	utils2 "github.com/CarrotVegeta/showstart/pkg/utils"
	jsoniter "github.com/json-iterator/go"
	"log"
	"math"
	"time"
)

const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateIndexOfChar(userID int) string {
	var n string
	r := make(map[int]string)
	i := make(map[string]int)
	for t := 0; t < len(str); t++ {
		n = string(str[t])
		r[t] = n
		i[n] = t
	}
	return encode(userID, r)
}
func encode(e int, n map[int]string) string {
	if e == 0 {
		return n[0]
	}
	var r string
	i := len(str)
	for e > 0 {
		r = n[e%i] + r
		e = int(math.Floor(float64(e / i)))
	}
	return r
}
func RequestWithBodyParam(action, url string, method string, query interface{}) (map[string]interface{}, error) {
	p := &models.Param{
		Action: action,
		Method: method,
		Query:  query,
		Body:   nil,
		Qtime:  time.Now().UnixNano() / 1e6,
		Ranstr: utils2.RandStr(8),
	}
	bs, _ := jsoniter.Marshal(p)
	c, err := utils2.AesECBEncrypt(bs, []byte(pkg.Secret))
	if err != nil {
		log.Fatal(err.Error())
	}
	d := base64.StdEncoding.EncodeToString(c)
	bp := GetBodyParam(d, string(bs))
	rm, err := request.HttpDo(url, bp)
	if err != nil {
		return nil, fmt.Errorf("发起post请求错误:%v", err.Error())
	}
	return rm, nil
}

func GetBodyParam(d, s string) *models.BodyParam {
	bp := &models.BodyParam{
		Data:     d,
		Appid:    pkg.APPID,
		Terminal: pkg.Terminal,
		Version:  pkg.Version,
	}
	bp.Sign = utils2.GeneMD5(s)
	return bp
}
