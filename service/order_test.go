package service

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	oc := &OrderConfig{
		GoodsId:           170092,
		SkuId:             "414e24781ce4f3c1e8aab424e19909a8",
		Price:             588,
		CommonPerformerID: 1225930,
		TelePhone:         "13118115155",
		SessionId:         3010537,
		StFlpv:            "Uri4jL8p8d4Ed1t55MI9",
	}
	m := Order(oc)
	fmt.Println(m)
}
