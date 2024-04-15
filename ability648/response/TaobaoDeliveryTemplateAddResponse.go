package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoDeliveryTemplateAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   模板对象
	*/
	DeliveryTemplate domain.TaobaoDeliveryTemplateAddDeliveryTemplate `json:"delivery_template,omitempty" `
}
