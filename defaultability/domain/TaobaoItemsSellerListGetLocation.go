package domain


type TaobaoItemsSellerListGetLocation struct {
    /*
        所在城市（中文名称）     */
    City  *string `json:"city,omitempty" `

    /*
        所在省份（中文名称）     */
    State  *string `json:"state,omitempty" `

}

func (s *TaobaoItemsSellerListGetLocation) SetCity(v string) *TaobaoItemsSellerListGetLocation {
    s.City = &v
    return s
}
func (s *TaobaoItemsSellerListGetLocation) SetState(v string) *TaobaoItemsSellerListGetLocation {
    s.State = &v
    return s
}
