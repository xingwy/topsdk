package domain


type AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest struct {
    /*
        子订单id     */
    SubTid  *string `json:"sub_tid,omitempty" `

    /*
        商品类型 0：普通品 1:赠品 2:成分品，默认0     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        成分品itemId     */
    CompItemId  *string `json:"comp_item_id,omitempty" `

    /*
        成分品skuId     */
    CompSkuId  *string `json:"comp_sku_id,omitempty" `

}

func (s *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest) SetSubTid(v string) *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest {
    s.SubTid = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest) SetItemType(v int64) *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest {
    s.ItemType = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest) SetCompItemId(v string) *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest {
    s.CompItemId = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest) SetCompSkuId(v string) *AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest {
    s.CompSkuId = &v
    return s
}
