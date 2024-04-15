package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type AlibabaAscpLogisticsSellerSendResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   -
	*/
	Result domain.AlibabaAscpLogisticsSellerSendResultDTO `json:"result,omitempty" `
}
