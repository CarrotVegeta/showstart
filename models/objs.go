package models

type GenerateOrder struct {
	Action string `json:"action"`
	Method string `json:"method"`
	Query  struct {
		OrderDetails      []*OrderDetails `json:"orderDetails"`
		CommonPerfomerIds []interface{}   `json:"commonPerfomerIds"`
		AreaCode          string          `json:"areaCode"`
		Telephone         string          `json:"telephone"`
		CustomerName      string          `json:"customerName"`
		ProvinceName      string          `json:"provinceName"`
		CityName          string          `json:"cityName"`
		Address           string          `json:"address"`
		TeamId            string          `json:"teamId"`
		CouponId          string          `json:"couponId"`
		CheckCode         string          `json:"checkCode"`
		Source            int             `json:"source"`
		Discount          int             `json:"discount"`
		SessionId         int             `json:"sessionId"`
		Freight           int             `json:"freight"`
		AmountPayable     string          `json:"amountPayable"`
		TotalAmount       int             `json:"totalAmount"`
		Partner           string          `json:"partner"`
		StFlpv            string          `json:"st_flpv"`
		TrackPath         string          `json:"trackPath"`
		Terminal          string          `json:"terminal"`
	} `json:"query"`
	Body   interface{} `json:"body"`
	Qtime  int64       `json:"qtime"`
	Ranstr string      `json:"ranstr"`
}
type OrderDetails struct {
	GoodsType int    `json:"goodsType"`
	SkuType   int    `json:"skuType"`
	Num       string `json:"num"`
	GoodsId   int    `json:"goodsId"`
	SkuId     string `json:"skuId"`
	//Price     string    `json:"price"`
}
type GetActivityDetails struct {
	ActivityId string `json:"activityId"`
	Coupon     string `json:"coupon"`
	StFlpv     string `json:"st_flpv"`
	Sign       string `json:"sign"`
	TrackPath  string `json:"trackPath"`
	Terminal   string `json:"terminal"`
}
type Param struct {
	Action string      `json:"action"`
	Method string      `json:"method"`
	Query  interface{} `json:"query"`
	Body   interface{} `json:"body"`
	Qtime  int64       `json:"qtime"`
	Ranstr string      `json:"ranstr"`
}
type BodyParam struct {
	Data     string `json:"data"`
	Sign     string `json:"sign"`
	Appid    string `json:"appid"`
	Terminal string `json:"terminal"`
	Version  string `json:"version"`
}
type TicketParam struct {
	ActivityId string `json:"activityId"`
	Coupon     string `json:"coupon"`
	StFlpv     string `json:"st_flpv"`
	Sign       string `json:"sign"`
	TrackPath  string `json:"trackPath"`
	Terminal   string `json:"terminal"`
}
type RequestSearch struct {
	PageNo      int    `json:"pageNo"`
	CityCode    string `json:"cityCode"`
	Keyword     string `json:"keyword"`
	Style       string `json:"style"`
	ActivityIds string `json:"activityIds"`
	CouponCode  string `json:"couponCode"`
	PerformerId string `json:"performerId"`
	HosterId    string `json:"hosterId"`
	SiteId      string `json:"siteId"`
	Tag         string `json:"tag"`
	TourId      string `json:"tourId"`
	ThemeId     string `json:"themeId"`
	StFlpv      string `json:"st_flpv"`
	Sign        string `json:"sign"`
	TrackPath   string `json:"trackPath"`
	Terminal    string `json:"terminal"`
}
type RequestSpectator struct {
	StFlpv    string `json:"st_flpv"`
	Sign      string `json:"sign"`
	TrackPath string `json:"trackPath"`
	Terminal  string `json:"terminal"`
}
