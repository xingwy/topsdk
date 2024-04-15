package domain


type TaobaoItemSellerGetLocation struct {
    /*
        所在城市（中文名称）     */
    City  *string `json:"city,omitempty" `

    /*
        所在省份（中文名称）     */
    State  *string `json:"state,omitempty" `

}

func (s *TaobaoItemSellerGetLocation) SetCity(v string) *TaobaoItemSellerGetLocation {
    s.City = &v
    return s
}
func (s *TaobaoItemSellerGetLocation) SetState(v string) *TaobaoItemSellerGetLocation {
    s.State = &v
    return s
}
