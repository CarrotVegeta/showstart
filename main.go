package main

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/service"
)

func main() {
	//m := service.GetActivityDetail()
	//fmt.Println(m)
	m := service.Order()
	fmt.Println(m)
}
