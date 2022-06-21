package main

import (
	"github.com/CarrotVegeta/showstart/server"
	"github.com/CarrotVegeta/showstart/server/router"
)

func main() {
	s := server.NewServer()
	s.RegisterRouter(&router.Activity{})
	s.RegisterRouter(&router.Ticket{})
	s.RegisterRouter(&router.OrderServer{})
	s.Start()
}
