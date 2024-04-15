package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoLogisticsDummySendResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回发货是否成功is_success
	*/
	Shipping domain.TaobaoLogisticsDummySendShipping `json:"shipping,omitempty" `
}
