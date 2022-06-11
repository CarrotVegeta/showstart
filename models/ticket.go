package models

type Ticket struct {
	SessionId    int    `json:"session_id"`
	Type         int    `json:"type"`
	TicketId     string `json:"ticket_id"`
	SellingPrice string `json:"selling_price"`
	GoodType     int    `json:"good_type"`
	TicketNum    int    `json:"ticket_num"`
	ActivityId   int    `json:"activity_id"`
	MemberNum    int    `json:"member_num"`
	RemainTicket int    `json:"remain_ticket"`
	TicketType   string `json:"ticket_type"`
	ShowTime     string `json:"show_time"`
}
type TicketModel *Ticket

func NewTicket() *Ticket {
	return &Ticket{}
}
