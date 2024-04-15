package domain


type TaobaoRefundDetailGetOperation struct {
    /*
        操作编码     */
    OperationCode  *string `json:"operation_code,omitempty" `

    /*
        操作提示文案     */
    Tips  *string `json:"tips,omitempty" `

}

func (s *TaobaoRefundDetailGetOperation) SetOperationCode(v string) *TaobaoRefundDetailGetOperation {
    s.OperationCode = &v
    return s
}
func (s *TaobaoRefundDetailGetOperation) SetTips(v string) *TaobaoRefundDetailGetOperation {
    s.Tips = &v
    return s
}
