package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoItemsSellerListGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   商品详细信息列表
	*/
	Items []domain.TaobaoItemsSellerListGetItem `json:"items,omitempty" `
}
