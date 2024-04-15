package domain


type TaobaoRefundDetailGetRefundDetail struct {
    /*
        退款版本号     */
    RefundVersion  *int64 `json:"refund_version,omitempty" `

    /*
        退款当前可以执行的操作     */
    AllowedOperations  *[]TaobaoRefundDetailGetOperation `json:"allowed_operations,omitempty" `

    /*
        退款当前不可以执行的操作     */
    NotAllowedOperations  *[]TaobaoRefundDetailGetOperation `json:"not_allowed_operations,omitempty" `

}

func (s *TaobaoRefundDetailGetRefundDetail) SetRefundVersion(v int64) *TaobaoRefundDetailGetRefundDetail {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoRefundDetailGetRefundDetail) SetAllowedOperations(v []TaobaoRefundDetailGetOperation) *TaobaoRefundDetailGetRefundDetail {
    s.AllowedOperations = &v
    return s
}
func (s *TaobaoRefundDetailGetRefundDetail) SetNotAllowedOperations(v []TaobaoRefundDetailGetOperation) *TaobaoRefundDetailGetRefundDetail {
    s.NotAllowedOperations = &v
    return s
}
