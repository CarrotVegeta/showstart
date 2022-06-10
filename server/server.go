package server

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/server/router"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Init("./config/config.yaml")
}

type GinServer interface {
	RegisterRouter(e *gin.Engine)
}
type Server struct {
	engine   *gin.Engine
	services []*GinServer
}

func (s *Server) RegisterRouter() {

}
func (s *Server) RegisterService() {
	s.services = append(s.services, &router.Activity{})
}
func (s *Server) Start() {
	s.engine.Run(fmt.Sprintf("0.0.0.0:%d", config.Conf.Port))
}
func NewServer() {
	s := &Server{}
	s.engine = gin.Default()
	s.Start()
}
