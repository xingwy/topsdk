package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type AlibabaAscpLogisticsSellerWriteoffResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   操作结果
	*/
	Result domain.AlibabaAscpLogisticsSellerWriteoffResultDTO `json:"result,omitempty" `
}
