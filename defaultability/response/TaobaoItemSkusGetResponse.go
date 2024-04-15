package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoItemSkusGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   Sku列表
	*/
	Skus []domain.TaobaoItemSkusGetSku `json:"skus,omitempty" `
}
