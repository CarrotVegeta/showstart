package models

type Order struct {
	ActivityID        int    `json:"activity_id"`
	TicketID          string `json:"ticket_id"`
	CommonPerformerID int    `json:"common_performer_id"`
	Telephone         string `json:"telephone"`
	CustomerName      string `json:"customer_name"`
	ProvinceName      string `json:"province_name"`
	CityName          string `json:"city_name"`
	Address           string `json:"address"`
	TicketNum         int    `json:"ticket_num"`
	GoodsType         int    `json:"goods_type"`
	SkuType           int    `json:"sku_type"`
	SellingPrice      string `json:"selling_price"`
	SessionId         int    `json:"session_id"`
}
type OrderModel *Order

func NewOrder() *Order {
	return &Order{}
}
