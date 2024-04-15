package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type AlibabaAscpLogisticsConsignResendResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   异步获取历史数据接口返回结果
	*/
	Result domain.AlibabaAscpLogisticsConsignResendResultDto `json:"result,omitempty" `
}
