package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TmallTraderateFeedsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回评价信息
	*/
	TmallRateInfo domain.TmallTraderateFeedsGetModel `json:"tmall_rate_info,omitempty" `
}
