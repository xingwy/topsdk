package domain


type TaobaoItemsSellerListGetSku struct {
    /*
        商品级别的条形码     */
    Barcode  *string `json:"barcode,omitempty" `

    /*
        sku创建日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *string `json:"created,omitempty" `

    /*
        sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        商家设置的外部id。天猫和集市的卖家，需要登录才能获取到自己的商家编码，不能获取到他人的商家编码。     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。     */
    Price  *string `json:"price,omitempty" `

    /*
        sku的销售属性组合字符串（颜色，大小，等等，可通过类目API获取某类目下的销售属性）,格式是p1:v1;p2:v2     */
    Properties  *string `json:"properties,omitempty" `

    /*
        sku所对应的销售属性的中文名字串，格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……     */
    PropertiesName  *string `json:"properties_name,omitempty" `

    /*
        属于这个sku的商品的数量，     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        skuFeature     */
    SkuFeature  *TaobaoItemsSellerListGetSkuFeature `json:"sku_feature,omitempty" `

}

func (s *TaobaoItemsSellerListGetSku) SetBarcode(v string) *TaobaoItemsSellerListGetSku {
    s.Barcode = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetCreated(v string) *TaobaoItemsSellerListGetSku {
    s.Created = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetModified(v string) *TaobaoItemsSellerListGetSku {
    s.Modified = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetOuterId(v string) *TaobaoItemsSellerListGetSku {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetPrice(v string) *TaobaoItemsSellerListGetSku {
    s.Price = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetProperties(v string) *TaobaoItemsSellerListGetSku {
    s.Properties = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetPropertiesName(v string) *TaobaoItemsSellerListGetSku {
    s.PropertiesName = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetQuantity(v int64) *TaobaoItemsSellerListGetSku {
    s.Quantity = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetSkuId(v int64) *TaobaoItemsSellerListGetSku {
    s.SkuId = &v
    return s
}
func (s *TaobaoItemsSellerListGetSku) SetSkuFeature(v TaobaoItemsSellerListGetSkuFeature) *TaobaoItemsSellerListGetSku {
    s.SkuFeature = &v
    return s
}
