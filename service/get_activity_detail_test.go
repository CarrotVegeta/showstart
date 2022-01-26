package service

import (
	"fmt"
	"testing"
)

func TestGetActivityDetail(t *testing.T) {
	res := GetActivityDetail("170309", "Uri4jL8p8d4Ed1t55MI9")
	fmt.Println(res)
}
