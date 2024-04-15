package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoOpencrmNodereportGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.TaobaoOpencrmNodereportGetNodeExecuteReportDto `json:"result,omitempty" `
}
