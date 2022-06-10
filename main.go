package main

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/server/router"
)

func main() {
	config.Init("./config/config.yaml")
	logger.Init()
	s := server.NewServer()
	s.RegisterRouter(&router.Activity{})
	s.Start()
}
