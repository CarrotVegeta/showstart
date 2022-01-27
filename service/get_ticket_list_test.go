package service

import (
	"fmt"
	"testing"
)

func TestGetTicketList(t *testing.T) {
	m := GetTicketList("170309")
	result := m["result"].(interface{})
	fmt.Println(result)
}
