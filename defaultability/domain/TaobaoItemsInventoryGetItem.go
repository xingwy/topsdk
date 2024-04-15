package domain

import (
        "topsdk/util"
    )

type TaobaoItemsInventoryGetItem struct {
    /*
        商品上传后的状态。onsale出售中，instock库中     */
    ApproveStatus  *string `json:"approve_status,omitempty" `

    /*
        商品id(注意：iid近期即将废弃，请用num_iid参数)     */
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
        有效期,7或者14（默认是14天）     */
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
        商家外部编码(可与商家外部系统对接)     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        是否在淘宝显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" `

    /*
        是否在外部网店显示     */
    IsEx  *bool `json:"is_ex,omitempty" `

    /*
        虚拟商品的状态字段     */
    IsVirtual  *bool `json:"is_virtual,omitempty" `

    /*
        商品首次上架时间     */
    FirstStartsTime  *util.LocalTime `json:"first_starts_time,omitempty" `

}

func (s *TaobaoItemsInventoryGetItem) SetApproveStatus(v string) *TaobaoItemsInventoryGetItem {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetIid(v string) *TaobaoItemsInventoryGetItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetNumIid(v int64) *TaobaoItemsInventoryGetItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetTitle(v string) *TaobaoItemsInventoryGetItem {
    s.Title = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetNick(v string) *TaobaoItemsInventoryGetItem {
    s.Nick = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetType(v string) *TaobaoItemsInventoryGetItem {
    s.Type = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetCid(v int64) *TaobaoItemsInventoryGetItem {
    s.Cid = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetSellerCids(v string) *TaobaoItemsInventoryGetItem {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetPicUrl(v string) *TaobaoItemsInventoryGetItem {
    s.PicUrl = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetNum(v int64) *TaobaoItemsInventoryGetItem {
    s.Num = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetProps(v string) *TaobaoItemsInventoryGetItem {
    s.Props = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetValidThru(v int64) *TaobaoItemsInventoryGetItem {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetListTime(v util.LocalTime) *TaobaoItemsInventoryGetItem {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetPrice(v string) *TaobaoItemsInventoryGetItem {
    s.Price = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetHasDiscount(v bool) *TaobaoItemsInventoryGetItem {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetHasInvoice(v bool) *TaobaoItemsInventoryGetItem {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetHasWarranty(v bool) *TaobaoItemsInventoryGetItem {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetHasShowcase(v bool) *TaobaoItemsInventoryGetItem {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetModified(v util.LocalTime) *TaobaoItemsInventoryGetItem {
    s.Modified = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetDelistTime(v util.LocalTime) *TaobaoItemsInventoryGetItem {
    s.DelistTime = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetPostageId(v int64) *TaobaoItemsInventoryGetItem {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetOuterId(v string) *TaobaoItemsInventoryGetItem {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetIsTaobao(v bool) *TaobaoItemsInventoryGetItem {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetIsEx(v bool) *TaobaoItemsInventoryGetItem {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetIsVirtual(v bool) *TaobaoItemsInventoryGetItem {
    s.IsVirtual = &v
    return s
}
func (s *TaobaoItemsInventoryGetItem) SetFirstStartsTime(v util.LocalTime) *TaobaoItemsInventoryGetItem {
    s.FirstStartsTime = &v
    return s
}
