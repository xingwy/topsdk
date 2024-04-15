package domain


type TaobaoTradeSimpleGetExpandCardInfo struct {
    /*
        买卡订单本金     */
    BasicPrice  *string `json:"basic_price,omitempty" `

    /*
        买卡订单权益金     */
    ExpandPrice  *string `json:"expand_price,omitempty" `

    /*
        用卡订单使用的本金     */
    BasicPriceUsed  *string `json:"basic_price_used,omitempty" `

    /*
        用卡订单使用的权益金     */
    ExpandPriceUsed  *string `json:"expand_price_used,omitempty" `

}

func (s *TaobaoTradeSimpleGetExpandCardInfo) SetBasicPrice(v string) *TaobaoTradeSimpleGetExpandCardInfo {
    s.BasicPrice = &v
    return s
}
func (s *TaobaoTradeSimpleGetExpandCardInfo) SetExpandPrice(v string) *TaobaoTradeSimpleGetExpandCardInfo {
    s.ExpandPrice = &v
    return s
}
func (s *TaobaoTradeSimpleGetExpandCardInfo) SetBasicPriceUsed(v string) *TaobaoTradeSimpleGetExpandCardInfo {
    s.BasicPriceUsed = &v
    return s
}
func (s *TaobaoTradeSimpleGetExpandCardInfo) SetExpandPriceUsed(v string) *TaobaoTradeSimpleGetExpandCardInfo {
    s.ExpandPriceUsed = &v
    return s
}
