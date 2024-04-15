package domain


type TaobaoRefundNegotiatereturnRenderReason struct {
    /*
        退款原因ID     */
    ReasonId  *int64 `json:"reason_id,omitempty" `

    /*
        退款原因文案     */
    ReasonText  *string `json:"reason_text,omitempty" `

}

func (s *TaobaoRefundNegotiatereturnRenderReason) SetReasonId(v int64) *TaobaoRefundNegotiatereturnRenderReason {
    s.ReasonId = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderReason) SetReasonText(v string) *TaobaoRefundNegotiatereturnRenderReason {
    s.ReasonText = &v
    return s
}
