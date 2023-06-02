package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyHandler func(c *gin.Context, reply *Reply)

func HandlerMiddleWare(h MyHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		reply := &Reply{}
		h(context, reply)
		if reply.Error != "" {
			context.AbortWithStatusJSON(http.StatusOK, &reply)
		}
	}
}
func Handler(h MyHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		reply := &Reply{}
		h(context, reply)
		context.AbortWithStatusJSON(http.StatusOK, &reply)
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
