package engine

import (
	"flag"
	"github.com/CarrotVegeta/showstart/config"
	"log"
)

type Server struct {
	ModM map[string]Engine
}

func (s *Server) RegisterMod(e Engine, mod string) {
	s.ModM[mod] = e
}
func (s *Server) RegisterServer(e Engine) {
	e.Open(s)
}
func (s *Server) Start() {
	args := flag.Args()
	if _, ok := s.ModM[args[0]]; !ok {
		log.Fatal("不存在的模式")
	}
	s.ModM[args[0]].Request()
}
func NewEngine() *Server {
	s := &Server{}
	s.ModM = make(map[string]Engine)
	return s
}
func init() {
	f := &Flag{}
	f.parseFlag()
	config.Init(f.ConfigPath)
}

type Flag struct {
	ConfigPath string
}

func (f *Flag) parseFlag() {
	configPath := flag.String("config", "", "config path")
	flag.Parse()
	f.ConfigPath = *configPath
}

type Engine interface {
	Request()
	Open(s *Server)
}
