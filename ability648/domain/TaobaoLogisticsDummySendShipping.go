package domain


type TaobaoLogisticsDummySendShipping struct {
    /*
        返回发货是否成功。     */
    IsSuccess  *bool `json:"is_success,omitempty" `

}

func (s *TaobaoLogisticsDummySendShipping) SetIsSuccess(v bool) *TaobaoLogisticsDummySendShipping {
    s.IsSuccess = &v
    return s
}
