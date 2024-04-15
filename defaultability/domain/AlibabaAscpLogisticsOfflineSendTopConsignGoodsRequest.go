package domain


type AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest struct {
    /*
        子订单id     */
    SubTid  *string `json:"sub_tid,omitempty" `

    /*
        成分品itemId     */
    CompItemId  *string `json:"comp_item_id,omitempty" `

    /*
        成分品skuId     */
    CompSkuId  *string `json:"comp_sku_id,omitempty" `

    /*
        品类型 0：标品/平台赠品，1：ERP线下赠品、2：成分品，默认为0     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        商品数量，不传默认为子单上的商品数量；支持不传，但不能传0或负值     */
    Num  *int64 `json:"num,omitempty" `

}

func (s *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) SetSubTid(v string) *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest {
    s.SubTid = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) SetCompItemId(v string) *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest {
    s.CompItemId = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) SetCompSkuId(v string) *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest {
    s.CompSkuId = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) SetItemType(v int64) *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest {
    s.ItemType = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) SetNum(v int64) *AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest {
    s.Num = &v
    return s
}
