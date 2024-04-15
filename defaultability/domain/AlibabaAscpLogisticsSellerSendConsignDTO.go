package domain


type AlibabaAscpLogisticsSellerSendConsignDTO struct {
    /*
        物流发货单号     */
    LpOrderId  *int64 `json:"lp_order_id,omitempty" `

    /*
        发货文案描述     */
    ConsignDesc  *string `json:"consign_desc,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerSendConsignDTO) SetLpOrderId(v int64) *AlibabaAscpLogisticsSellerSendConsignDTO {
    s.LpOrderId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendConsignDTO) SetConsignDesc(v string) *AlibabaAscpLogisticsSellerSendConsignDTO {
    s.ConsignDesc = &v
    return s
}
