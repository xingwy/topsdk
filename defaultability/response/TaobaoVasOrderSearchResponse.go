package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoVasOrderSearchResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   商品订单对象
	*/
	ArticleBizOrders []domain.TaobaoVasOrderSearchArticleBizOrder `json:"article_biz_orders,omitempty" `
	/*
	   总记录数
	*/
	TotalItem int64 `json:"total_item,omitempty" `
}
