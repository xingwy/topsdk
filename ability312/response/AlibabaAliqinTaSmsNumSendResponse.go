package response

import (
	"github.com/xingwy/topsdk/ability312/domain"
)

type AlibabaAliqinTaSmsNumSendResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Result domain.AlibabaAliqinTaSmsNumSendBizResult `json:"result,omitempty" `
}
