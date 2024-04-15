package domain


type TaobaoTradeSellerflagsGetSellerFlag struct {
    /*
        卖家ID     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        插旗ID     */
    FlagId  *int64 `json:"flag_id,omitempty" `

    /*
        插旗备注     */
    TagContent  *string `json:"tag_content,omitempty" `

}

func (s *TaobaoTradeSellerflagsGetSellerFlag) SetSellerId(v int64) *TaobaoTradeSellerflagsGetSellerFlag {
    s.SellerId = &v
    return s
}
func (s *TaobaoTradeSellerflagsGetSellerFlag) SetFlagId(v int64) *TaobaoTradeSellerflagsGetSellerFlag {
    s.FlagId = &v
    return s
}
func (s *TaobaoTradeSellerflagsGetSellerFlag) SetTagContent(v string) *TaobaoTradeSellerflagsGetSellerFlag {
    s.TagContent = &v
    return s
}
