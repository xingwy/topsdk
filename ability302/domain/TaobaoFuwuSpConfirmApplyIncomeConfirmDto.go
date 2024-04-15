package domain


type TaobaoFuwuSpConfirmApplyIncomeConfirmDto struct {
    /*
        appkey     */
    Appkey  *string `json:"appkey,omitempty" `

    /*
        确认扩展信息     */
    Extend  *string `json:"extend,omitempty" `

    /*
        确认金额     */
    Fee  *int64 `json:"fee,omitempty" `

    /*
        卖家nick     */
    Nick  *string `json:"nick,omitempty" `

    /*
        服务市场有效订单ID     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        外部确认账单ID     */
    OutConfirmId  *string `json:"out_confirm_id,omitempty" `

    /*
        外部订单ID     */
    OutOrderId  *string `json:"out_order_id,omitempty" `

}

func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetAppkey(v string) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.Appkey = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetExtend(v string) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.Extend = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetFee(v int64) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.Fee = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetNick(v string) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.Nick = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetOrderId(v int64) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.OrderId = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetOutConfirmId(v string) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.OutConfirmId = &v
    return s
}
func (s *TaobaoFuwuSpConfirmApplyIncomeConfirmDto) SetOutOrderId(v string) *TaobaoFuwuSpConfirmApplyIncomeConfirmDto {
    s.OutOrderId = &v
    return s
}
