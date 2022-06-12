package server

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinServer interface {
	RegisterRouter(e *gin.Engine)
}
type Server struct {
	engine *gin.Engine
}

func (s *Server) RegisterRouter(g GinServer) {
	g.RegisterRouter(s.engine)
}
func (s *Server) Start() {
	logger.FileLog.Info("server start 0.0.0.0:%d", config.Conf.Server.Port)
	err := s.engine.Run(fmt.Sprintf("0.0.0.0:%d", config.Conf.Server.Port))
	if err != nil {
		logger.FileLog.Error(err.Error())
		return
	}
}

func NewServer() *Server {
	s := &Server{}
	s.engine = gin.Default()
	s.engine.Use(CorsMiddleware())
	s.engine.Use(HandlerMiddleWare()(SetUserInfo))
	err := s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err.Error())
	}
	return s
}

var User *models.User

func SetUserInfo(c *gin.Context) (reply *Reply) {
	reply = NewReply()
	User = models.NewUser()
	User.CUUSERREF = c.GetHeader("Cuuserref")
	if User.CUUSERREF == "" {
		reply.Error = "cuuseref is null"
		return
	}
	User.StFlpv = c.GetHeader("St_flpv")
	if User.StFlpv == "" {
		reply.Error = "st_flpv is null"
		return
	}
	User.CUSUT = c.GetHeader("Cusut")
	if User.CUSUT == "" {
		reply.Error = "cusut is null"
		return
	}
	return
}

//CorsMiddleware 处理跨域请求,支持options访问
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,Cuuserref,St_flpv,Cusut")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
