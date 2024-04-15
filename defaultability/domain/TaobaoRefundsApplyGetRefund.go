package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundsApplyGetRefund struct {
	/*
	   退款单号     */
	RefundId *string `json:"refund_id,omitempty" `

	/*
	   淘宝交易单号     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   卖家昵称     */
	SellerNick *string `json:"seller_nick,omitempty" `

	/*
	   交易总金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	TotalFee *string `json:"total_fee,omitempty" `

	/*
	   退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
	Status *string `json:"status,omitempty" `

	/*
	   退款申请时间。格式:yyyy-MM-dd HH:mm:ss     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	RefundFee *string `json:"refund_fee,omitempty" `

	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   买家openUid     */
	BuyerOpenUid *string `json:"buyer_open_uid,omitempty" `

	/*
	   买家昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" `
}

func (s *TaobaoRefundsApplyGetRefund) SetRefundId(v string) *TaobaoRefundsApplyGetRefund {
	s.RefundId = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetTid(v int64) *TaobaoRefundsApplyGetRefund {
	s.Tid = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetSellerNick(v string) *TaobaoRefundsApplyGetRefund {
	s.SellerNick = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetTotalFee(v string) *TaobaoRefundsApplyGetRefund {
	s.TotalFee = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetStatus(v string) *TaobaoRefundsApplyGetRefund {
	s.Status = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetCreated(v util.LocalTime) *TaobaoRefundsApplyGetRefund {
	s.Created = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetRefundFee(v string) *TaobaoRefundsApplyGetRefund {
	s.RefundFee = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetTitle(v string) *TaobaoRefundsApplyGetRefund {
	s.Title = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetBuyerOpenUid(v string) *TaobaoRefundsApplyGetRefund {
	s.BuyerOpenUid = &v
	return s
}
func (s *TaobaoRefundsApplyGetRefund) SetBuyerNick(v string) *TaobaoRefundsApplyGetRefund {
	s.BuyerNick = &v
	return s
}
