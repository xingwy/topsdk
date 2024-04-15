package domain


type TaobaoRefundStatusGetQueryRefundStatusResponse struct {
    /*
        退款单id     */
    RefundId  *int64 `json:"refund_id,omitempty" `

    /*
        淘宝交易单号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        子订单号     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        更新时间。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoRefundStatusGetQueryRefundStatusResponse) SetRefundId(v int64) *TaobaoRefundStatusGetQueryRefundStatusResponse {
    s.RefundId = &v
    return s
}
func (s *TaobaoRefundStatusGetQueryRefundStatusResponse) SetTid(v int64) *TaobaoRefundStatusGetQueryRefundStatusResponse {
    s.Tid = &v
    return s
}
func (s *TaobaoRefundStatusGetQueryRefundStatusResponse) SetOid(v int64) *TaobaoRefundStatusGetQueryRefundStatusResponse {
    s.Oid = &v
    return s
}
func (s *TaobaoRefundStatusGetQueryRefundStatusResponse) SetModified(v string) *TaobaoRefundStatusGetQueryRefundStatusResponse {
    s.Modified = &v
    return s
}
func (s *TaobaoRefundStatusGetQueryRefundStatusResponse) SetStatus(v string) *TaobaoRefundStatusGetQueryRefundStatusResponse {
    s.Status = &v
    return s
}
