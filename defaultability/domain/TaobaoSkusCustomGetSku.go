package domain


type TaobaoSkusCustomGetSku struct {
    /*
        sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        sku所属商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

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
        属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。     */
    Price  *string `json:"price,omitempty" `

    /*
        商家设置的外部id。天猫和集市的卖家，需要登录才能获取到自己的商家编码，不能获取到他人的商家编码。     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        sku创建日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *string `json:"created,omitempty" `

    /*
        sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        sku状态。 normal:正常 ；delete:删除     */
    Status  *string `json:"status,omitempty" `

    /*
        sku所属商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

}

func (s *TaobaoSkusCustomGetSku) SetSkuId(v int64) *TaobaoSkusCustomGetSku {
    s.SkuId = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetIid(v string) *TaobaoSkusCustomGetSku {
    s.Iid = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetProperties(v string) *TaobaoSkusCustomGetSku {
    s.Properties = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetPropertiesName(v string) *TaobaoSkusCustomGetSku {
    s.PropertiesName = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetQuantity(v int64) *TaobaoSkusCustomGetSku {
    s.Quantity = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetPrice(v string) *TaobaoSkusCustomGetSku {
    s.Price = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetOuterId(v string) *TaobaoSkusCustomGetSku {
    s.OuterId = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetCreated(v string) *TaobaoSkusCustomGetSku {
    s.Created = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetModified(v string) *TaobaoSkusCustomGetSku {
    s.Modified = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetStatus(v string) *TaobaoSkusCustomGetSku {
    s.Status = &v
    return s
}
func (s *TaobaoSkusCustomGetSku) SetNumIid(v int64) *TaobaoSkusCustomGetSku {
    s.NumIid = &v
    return s
}
