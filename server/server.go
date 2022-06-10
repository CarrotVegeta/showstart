package server

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
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
	err := s.engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err.Error())
	}
	return s
}
