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
	CronTime          string `json:"cron_time"`
	OrderNum          int    `json:"order_num"`
	TicketNum         int    `json:"ticket_num"`
	Ticker            int64  `json:"ticker"`
}
type OrderModel *Order

func NewOrder() *Order {
	return &Order{}
}
