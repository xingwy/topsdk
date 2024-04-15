package response

import (
	"github.com/xingwy/topsdk/ability200/domain"
)

type TaobaoRefundNegotiatereturnRenderResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   申请协商提示文案
	*/
	ApplyTips string `json:"apply_tips,omitempty" `
	/*
	   退款版本号
	*/
	RefundVersion int64 `json:"refund_version,omitempty" `
	/*
	   可以协商的退款原因列表
	*/
	ReasonList []domain.TaobaoRefundNegotiatereturnRenderReason `json:"reason_list,omitempty" `
	/*
	   可以协商的最大退款金额
	*/
	MaxRefundFee domain.TaobaoRefundNegotiatereturnRenderMaxRefundFee `json:"max_refund_fee,omitempty" `
	/*
	   卖家退货地址列表
	*/
	AddressList []domain.TaobaoRefundNegotiatereturnRenderAddress `json:"address_list,omitempty" `
}
