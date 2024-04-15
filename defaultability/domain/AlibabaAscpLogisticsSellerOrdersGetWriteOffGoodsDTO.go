package domain


type AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO struct {
    /*
        商品Id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        数量     */
    GoodsQuantity  *int64 `json:"goods_quantity,omitempty" `

    /*
        名称     */
    GoodsName  *string `json:"goods_name,omitempty" `

    /*
        商品sku id     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        商品图片地址     */
    GoodsPicId  *string `json:"goods_pic_id,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetItemId(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.ItemId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetPrice(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.Price = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetGoodsQuantity(v int64) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.GoodsQuantity = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetGoodsName(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.GoodsName = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetSkuId(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.SkuId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO) SetGoodsPicId(v string) *AlibabaAscpLogisticsSellerOrdersGetWriteOffGoodsDTO {
    s.GoodsPicId = &v
    return s
}
