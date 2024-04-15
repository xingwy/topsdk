package domain


type AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO struct {
    /*
        交易单所包含的商品列表     */
    GoodsList  *[]AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO `json:"goods_list,omitempty" `

    /*
        核销订单Id     */
    LpOrderId  *string `json:"lp_order_id,omitempty" `

    /*
        淘宝交易Id     */
    TradeId  *string `json:"trade_id,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO) SetGoodsList(v []AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO {
    s.GoodsList = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO) SetLpOrderId(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO {
    s.LpOrderId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO) SetTradeId(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffOrderDTO {
    s.TradeId = &v
    return s
}
