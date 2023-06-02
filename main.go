package main

import (
	"github.com/CarrotVegeta/showstart/api"
	"github.com/CarrotVegeta/showstart/api/router"
)

func main() {
	s := api.NewServer()
	s.RegisterRouter(&router.Activity{})
	s.RegisterRouter(&router.Ticket{})
	s.RegisterRouter(&router.OrderServer{})
	s.RegisterRouter(&router.SpectatorServer{})
	s.Start()
}
