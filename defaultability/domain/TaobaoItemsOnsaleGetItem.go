package domain

import (
        "topsdk/util"
    )

type TaobaoItemsOnsaleGetItem struct {
    /*
        商品上传后的状态。onsale出售中，instock库中     */
    ApproveStatus  *string `json:"approve_status,omitempty" `

    /*
        商品iid     */
    Iid  *string `json:"iid,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品标题,不能超过60字节     */
    Title  *string `json:"title,omitempty" `

    /*
        卖家昵称     */
    Nick  *string `json:"nick,omitempty" `

    /*
        商品类型(fixed:一口价;auction:拍卖)注：取消团购     */
    Type  *string `json:"type,omitempty" `

    /*
        商品所属的叶子类目 id     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        商品所属的店铺内卖家自定义类目列表     */
    SellerCids  *string `json:"seller_cids,omitempty" `

    /*
        商品主图片地址     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        商品数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        商品属性 格式：pid:vid;pid:vid     */
    Props  *string `json:"props,omitempty" `

    /*
        有效期,7或者14（默认是7天）     */
    ValidThru  *int64 `json:"valid_thru,omitempty" `

    /*
        上架时间（格式：yyyy-MM-dd HH:mm:ss）     */
    ListTime  *util.LocalTime `json:"list_time,omitempty" `

    /*
        商品价格，格式：5.00；单位：元；精确到：分     */
    Price  *string `json:"price,omitempty" `

    /*
        支持会员打折,true/false     */
    HasDiscount  *bool `json:"has_discount,omitempty" `

    /*
        是否有发票,true/false     */
    HasInvoice  *bool `json:"has_invoice,omitempty" `

    /*
        是否有保修,true/false     */
    HasWarranty  *bool `json:"has_warranty,omitempty" `

    /*
        橱窗推荐,true/false     */
    HasShowcase  *bool `json:"has_showcase,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        下架时间（格式：yyyy-MM-dd HH:mm:ss）     */
    DelistTime  *util.LocalTime `json:"delist_time,omitempty" `

    /*
        宝贝所属的运费模板ID，如果没有返回则说明没有使用运费模板     */
    PostageId  *int64 `json:"postage_id,omitempty" `

    /*
        商家外部编码(可与商家外部系统对接)。需要授权才能获取。     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        是否在外部网店显示     */
    IsEx  *bool `json:"is_ex,omitempty" `

    /*
        虚拟商品的状态字段     */
    IsVirtual  *bool `json:"is_virtual,omitempty" `

    /*
        是否在淘宝显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" `

    /*
        商品销量     */
    SoldQuantity  *int64 `json:"sold_quantity,omitempty" `

    /*
        是否为达尔文挂接成功了的商品     */
    IsCspu  *bool `json:"is_cspu,omitempty" `

    /*
        商品首次上架时间     */
    FirstStartsTime  *util.LocalTime `json:"first_starts_time,omitempty" `

}

func (s *TaobaoItemsOnsaleGetItem) SetApproveStatus(v string) *TaobaoItemsOnsaleGetItem {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetIid(v string) *TaobaoItemsOnsaleGetItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetNumIid(v int64) *TaobaoItemsOnsaleGetItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetTitle(v string) *TaobaoItemsOnsaleGetItem {
    s.Title = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetNick(v string) *TaobaoItemsOnsaleGetItem {
    s.Nick = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetType(v string) *TaobaoItemsOnsaleGetItem {
    s.Type = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetCid(v int64) *TaobaoItemsOnsaleGetItem {
    s.Cid = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetSellerCids(v string) *TaobaoItemsOnsaleGetItem {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetPicUrl(v string) *TaobaoItemsOnsaleGetItem {
    s.PicUrl = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetNum(v int64) *TaobaoItemsOnsaleGetItem {
    s.Num = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetProps(v string) *TaobaoItemsOnsaleGetItem {
    s.Props = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetValidThru(v int64) *TaobaoItemsOnsaleGetItem {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetListTime(v util.LocalTime) *TaobaoItemsOnsaleGetItem {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetPrice(v string) *TaobaoItemsOnsaleGetItem {
    s.Price = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetHasDiscount(v bool) *TaobaoItemsOnsaleGetItem {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetHasInvoice(v bool) *TaobaoItemsOnsaleGetItem {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetHasWarranty(v bool) *TaobaoItemsOnsaleGetItem {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetHasShowcase(v bool) *TaobaoItemsOnsaleGetItem {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetModified(v util.LocalTime) *TaobaoItemsOnsaleGetItem {
    s.Modified = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetDelistTime(v util.LocalTime) *TaobaoItemsOnsaleGetItem {
    s.DelistTime = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetPostageId(v int64) *TaobaoItemsOnsaleGetItem {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetOuterId(v string) *TaobaoItemsOnsaleGetItem {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetIsEx(v bool) *TaobaoItemsOnsaleGetItem {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetIsVirtual(v bool) *TaobaoItemsOnsaleGetItem {
    s.IsVirtual = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetIsTaobao(v bool) *TaobaoItemsOnsaleGetItem {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetSoldQuantity(v int64) *TaobaoItemsOnsaleGetItem {
    s.SoldQuantity = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetIsCspu(v bool) *TaobaoItemsOnsaleGetItem {
    s.IsCspu = &v
    return s
}
func (s *TaobaoItemsOnsaleGetItem) SetFirstStartsTime(v util.LocalTime) *TaobaoItemsOnsaleGetItem {
    s.FirstStartsTime = &v
    return s
}
