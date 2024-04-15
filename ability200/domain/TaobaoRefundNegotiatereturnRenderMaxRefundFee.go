package domain


type TaobaoRefundNegotiatereturnRenderMaxRefundFee struct {
    /*
        可以协商的最大退款金额     */
    MaxRefundFee  *int64 `json:"max_refund_fee,omitempty" `

}

func (s *TaobaoRefundNegotiatereturnRenderMaxRefundFee) SetMaxRefundFee(v int64) *TaobaoRefundNegotiatereturnRenderMaxRefundFee {
    s.MaxRefundFee = &v
    return s
}
