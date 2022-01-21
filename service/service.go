package service

import (
	"encoding/base64"
	"github.com/CarrotVegeta/showstart/request"
	"github.com/CarrotVegeta/showstart/utils"
	"io/ioutil"
	"log"
	"strings"
)

func GetActivityDetail() map[string]interface{} {
	m := make(map[string]interface{})
	m["data"] = "RdKnoWKXGxZP0zYw1ohcYvZbIyor3SpU3DeN8H53n1LW5PjbQdqFCFCiM8cBUfFO4clLwzs+dME5t5ajE68fEj+I0EmGeoCrC5bNRlwR4xXmzY0QM8ki0G3S/1dnd1zwIxWL/h5Fx1/P3I8durNbfu5qSMilOkrMqM9M4W6m/M0FEq0rssxosRykT7vCKolDgaIU93Tl980WUUtJ89bfCN78ABBOjEl+5942Xtc2gTGOFd3jLz2B+o0tM1p5G9w4KUBrnPtP3NMDl+fzXuga9S2FlU2301ILeKk8Lab/gz1Cnq4DMZduIbfI1NV2zhsEhU4Kv3sA2TWxXmT6iNFk7n7YCPqTB6TGS0bpA+Ltlzo="
	m["sign"] = "1e45f745caa4989f140ff9d296abcaef"
	m["appid"] = "wap"
	m["terminal"] = "wap"
	m["version"] = "997"
	m, err := request.HttpDo(request.ActivityDetail, "4ea93e5909142dc64f673ed009134cbb", m)
	if err != nil {
		log.Fatal(err.Error())
	}
	return m
}
func Order() map[string]interface{} {
	m := make(map[string]interface{})
	f, err := ioutil.ReadFile("./123.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	s := string(f)
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	b := `0RGF99CtUajPF0Ny`
	c, err := utils.AesECBEncrypt([]byte(s), []byte(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	d := base64.StdEncoding.EncodeToString(c)
	m["data"] = d
	m["sign"] = utils.GeneMD5(s)
	m["appid"] = "wap"
	m["terminal"] = "wap"
	m["version"] = "997"
	e, err := request.HttpDo(request.Order, "4ea93e5909142dc64f673ed009134cbb", m)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
