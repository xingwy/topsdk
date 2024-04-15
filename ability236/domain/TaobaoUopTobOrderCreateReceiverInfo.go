package domain


type TaobaoUopTobOrderCreateReceiverInfo struct {
    /*
        收货人移动电话     */
    Mobile  *string `json:"mobile,omitempty" `

    /*
        收货人姓名     */
    Name  *string `json:"name,omitempty" `

    /*
        收货人详细地址     */
    DetailAddress  *string `json:"detail_address,omitempty" `

    /*
        收货人镇     */
    Town  *string `json:"town,omitempty" `

    /*
        收货人区     */
    Area  *string `json:"area,omitempty" `

    /*
        收货人市     */
    City  *string `json:"city,omitempty" `

    /*
        收货人省     */
    Province  *string `json:"province,omitempty" `

}

func (s *TaobaoUopTobOrderCreateReceiverInfo) SetMobile(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.Mobile = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetName(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.Name = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetDetailAddress(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.DetailAddress = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetTown(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.Town = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetArea(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.Area = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetCity(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.City = &v
    return s
}
func (s *TaobaoUopTobOrderCreateReceiverInfo) SetProvince(v string) *TaobaoUopTobOrderCreateReceiverInfo {
    s.Province = &v
    return s
}
