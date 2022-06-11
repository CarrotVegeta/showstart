package server

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/gin-gonic/gin"
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
