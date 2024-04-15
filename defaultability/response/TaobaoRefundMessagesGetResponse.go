package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoRefundMessagesGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询到的退款留言/凭证列表
	*/
	RefundMessages []domain.TaobaoRefundMessagesGetRefundMessage `json:"refund_messages,omitempty" `
	/*
	   查询到的退款留言/凭证总数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
}
