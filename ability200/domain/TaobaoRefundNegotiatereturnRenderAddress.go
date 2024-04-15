package domain


type TaobaoRefundNegotiatereturnRenderAddress struct {
    /*
        地址ID     */
    AddressId  *int64 `json:"address_id,omitempty" `

    /*
        收件人姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        邮政编码     */
    PostCode  *string `json:"post_code,omitempty" `

    /*
        收件人手机     */
    Mobile  *string `json:"mobile,omitempty" `

    /*
        国家     */
    CountryName  *string `json:"country_name,omitempty" `

    /*
        省     */
    ProvinceName  *string `json:"province_name,omitempty" `

    /*
        市     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        区     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        乡镇/街道     */
    TownName  *string `json:"town_name,omitempty" `

    /*
        乡镇/街道地址详情     */
    AddressDetail  *string `json:"address_detail,omitempty" `

    /*
        行政区划代码     */
    DivisionCode  *string `json:"division_code,omitempty" `

}

func (s *TaobaoRefundNegotiatereturnRenderAddress) SetAddressId(v int64) *TaobaoRefundNegotiatereturnRenderAddress {
    s.AddressId = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetReceiverName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetPostCode(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.PostCode = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetMobile(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.Mobile = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetCountryName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.CountryName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetProvinceName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.ProvinceName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetCityName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.CityName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetAreaName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.AreaName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetTownName(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.TownName = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetAddressDetail(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.AddressDetail = &v
    return s
}
func (s *TaobaoRefundNegotiatereturnRenderAddress) SetDivisionCode(v string) *TaobaoRefundNegotiatereturnRenderAddress {
    s.DivisionCode = &v
    return s
}
