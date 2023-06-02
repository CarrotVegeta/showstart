package service

import (
	"github.com/CarrotVegeta/showstart/api"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/pkg"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Spectator struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	ShowDocumentNumber string `json:"showDocumentNumber"`
}

func GetSpectators(c *gin.Context, reply *api.Reply) {
	res, err := GetSpectatorList()
	if err != nil {
		logger.FileLog.Error(err.Error())
		reply.Error = err.Error()
		return
	}
	result := res["result"].([]interface{})
	bs, _ := jsoniter.Marshal(result)
	var sps []Spectator
	err = jsoniter.Unmarshal(bs, &sps)
	if err != nil {
		logger.FileLog.Error(err.Error())
		reply.Error = err.Error()
		return
	}
	reply.Data = sps
}
func GetSpectatorList() (map[string]interface{}, error) {
	return RequestWithBodyParam(pkg.GetSpectatorAction, pkg.HW, http.MethodPost, nil)
}
