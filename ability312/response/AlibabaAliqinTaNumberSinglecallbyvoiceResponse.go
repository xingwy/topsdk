package response

import (
	"github.com/xingwy/topsdk/ability312/domain"
)

type AlibabaAliqinTaNumberSinglecallbyvoiceResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.AlibabaAliqinTaNumberSinglecallbyvoiceResult `json:"result,omitempty" `
}
