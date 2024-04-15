package domain


type TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto struct {
    /*
        是否执行成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        订单号     */
    OrderCode  *string `json:"order_code,omitempty" `

}

func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto) SetSuccess(v bool) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto {
    s.Success = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto) SetOrderCode(v string) *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto {
    s.OrderCode = &v
    return s
}
