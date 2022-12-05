package main

import (
	"github.com/CarrotVegeta/showstart/api"
	router2 "github.com/CarrotVegeta/showstart/api/router"
)

func main() {
	s := api.NewServer()
	s.RegisterRouter(&router2.Activity{})
	s.RegisterRouter(&router2.Ticket{})
	s.RegisterRouter(&router2.OrderServer{})
	s.Start()
}
