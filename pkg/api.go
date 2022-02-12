package pkg

const (
	BaseUrl         = "https://wap.showstart.com/api"
	Detail          = "/hw/000IiV00ggre"    //获取详细信息
	GetList         = "/hw/00000000ggre"    //获取列表
	Order           = "/appnj/02cCBf00ggre" //普通下单
	Order2          = "/appnj/01cDb300ggre" //带观演人下单
	CoreOrderResult = "/appnj/02cDGH00ggre"
)

var orderSkuTypeM = map[int]string{
	1: Order,
	2: Order2,
}

func GetOrderUrl(skuType int) string {
	return orderSkuTypeM[skuType]
}
