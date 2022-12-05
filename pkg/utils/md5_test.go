package utils

import (
	"fmt"
	"testing"
)

func TestGeneMD5(t *testing.T) {
	a := GeneMD5(`{"action":"/wap/activity/details","method":"POST","query":{"activityId":"170309","coupon":"","st_flpv":"Uri4jL8p8d4Ed1t55MI9","sign":"611344d559eebc9fd940dee43e778f4c","trackPath":"","terminal":"wap"},"body":null,"qtime":1643261858468,"ranstr":"ePnQ7H48"}`)
	fmt.Println(a)
}
