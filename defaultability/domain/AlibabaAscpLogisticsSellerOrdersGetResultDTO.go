package domain


type AlibabaAscpLogisticsSellerOrdersGetResultDTO struct {
    /*
        返回核销订单列表     */
    WriteoffOrderList  *[]AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO `json:"writeoff_order_list,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerOrdersGetResultDTO) SetWriteoffOrderList(v []AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO) *AlibabaAscpLogisticsSellerOrdersGetResultDTO {
    s.WriteoffOrderList = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsSellerOrdersGetResultDTO {
    s.Success = &v
    return s
}
