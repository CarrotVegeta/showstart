package utils

import (
	"fmt"
	"testing"
)

func TestGeneMD5(t *testing.T) {
	a := GeneMD5(`{"action":"/order/order","method":"POST","query":{"orderDetails":[{"goodsType":1,"skuType":1,"num":"1","goodsId":169641,"skuId":"16724ccc8e49e61a4740a99cb700149f","price":140}],"commonPerfomerIds":[],"areaCode":"","telephone":"13118115155","customerName":"","provinceName":"","cityName":"","address":"","teamId":"","couponId":"","checkCode":"","source":0,"discount":0,"sessionId":3008317,"freight":0,"amountPayable":"140.00","totalAmount":120,"partner":"","st_flpv":"Uri4jL8p8d4Ed1t55MI9","sign":"611344d559eebc9fd940dee43e778f4c","trackPath":"","terminal":"wap"},"body":null,"qtime":1642744722756,"ranstr":"RS6pP4R6"}`)
	fmt.Println(a)
}
