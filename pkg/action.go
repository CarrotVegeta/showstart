package pkg

import utils2 "github.com/CarrotVegeta/showstart/pkg/utils"

const (
	GetActivityDetailsAction = "/wap/activity/details"
	OrderAction              = "/order/order"
	GetTicketListAction      = "/wap/activity/V2/ticket/list"
	GetActivityListAction    = "/wap/activity/list"
	GetSpectatorAction       = "/wap/cp/list"
	GetCoreOrderResultAction = "/order/getCoreOrderResult"
	WapUserLogin             = "/wap/user/login"
)

func GenerateUrl(u string, userID, eventID int) string {
	u += utils2.GenerateUrlApi(3876440, 196660, 0)
	return u
}
