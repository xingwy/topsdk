package response

import (
	"github.com/xingwy/topsdk/ability236/domain"
)

type TaobaoLogisticsExpressModifyAppointResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   调用结果
	*/
	Result domain.TaobaoLogisticsExpressModifyAppointSingleResultDto `json:"result,omitempty" `
}
