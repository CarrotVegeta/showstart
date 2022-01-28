package main

import (
	"github.com/CarrotVegeta/showstart/engine"
	"github.com/CarrotVegeta/showstart/mod"
)

func main() {
	s := engine.NewEngine()
	s.RegisterServer(&mod.ActivityList{})
	s.RegisterServer(&mod.ActivityDetail{})
	s.RegisterServer(&mod.Spectator{})
	s.RegisterServer(&mod.Order{})
	s.RegisterServer(&mod.Ticket{})
	s.Start()
}
