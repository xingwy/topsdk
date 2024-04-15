package response

import (
	"github.com/xingwy/topsdk/ability236/domain"
)

type TaobaoUopTobOrderCreateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   flag
	*/
	Flag string `json:"flag,omitempty" `
	/*
	   message
	*/
	Message string `json:"message,omitempty" `
	/*
	   订单
	*/
	DeliveryOrders []domain.TaobaoUopTobOrderCreateDeliveryorder `json:"delivery_orders,omitempty" `
}
