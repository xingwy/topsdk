package domain


type TaobaoRefundGetCombineSubItemDO struct {
    /*
        商品数字编号      */
    NumIid  *string `json:"num_iid,omitempty" `

    /*
        商品标题     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        sku_id     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        商家外部编码(可与商家外部系统对接)     */
    OuterIid  *string `json:"outer_iid,omitempty" `

    /*
        123     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

}

func (s *TaobaoRefundGetCombineSubItemDO) SetNumIid(v string) *TaobaoRefundGetCombineSubItemDO {
    s.NumIid = &v
    return s
}
func (s *TaobaoRefundGetCombineSubItemDO) SetItemName(v string) *TaobaoRefundGetCombineSubItemDO {
    s.ItemName = &v
    return s
}
func (s *TaobaoRefundGetCombineSubItemDO) SetQuantity(v int64) *TaobaoRefundGetCombineSubItemDO {
    s.Quantity = &v
    return s
}
func (s *TaobaoRefundGetCombineSubItemDO) SetSkuId(v string) *TaobaoRefundGetCombineSubItemDO {
    s.SkuId = &v
    return s
}
func (s *TaobaoRefundGetCombineSubItemDO) SetOuterIid(v string) *TaobaoRefundGetCombineSubItemDO {
    s.OuterIid = &v
    return s
}
func (s *TaobaoRefundGetCombineSubItemDO) SetOuterSkuId(v string) *TaobaoRefundGetCombineSubItemDO {
    s.OuterSkuId = &v
    return s
}
