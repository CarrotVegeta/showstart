package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyHandler func(c *gin.Context) (reply *Reply)

func Handler() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(context *gin.Context) {
			reply := h(context)
			context.AbortWithStatusJSON(http.StatusOK, &reply)
		}
	}
}

type Reply struct {
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

func NewReply() *Reply {
	return &Reply{}
}

func (r *Reply) WithData() {

}
