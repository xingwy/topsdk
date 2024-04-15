package request

import (
        "topsdk/util"
    )

type TaobaoItemsOnsaleGetRequest struct {
    /*
        需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        搜索字段。搜索商品的title。     */
    Q  *string `json:"q,omitempty" required:"false" `
    /*
        商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到     */
    Cid  *int64 `json:"cid,omitempty" required:"false" `
    /*
        卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)     */
    SellerCids  *string `json:"seller_cids,omitempty" required:"false" `
    /*
        页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc defalutValue��list_time:desc    */
    OrderBy  *string `json:"order_by,omitempty" required:"false" `
    /*
        是否参与会员折扣。可选值：true，false。默认不过滤该条件     */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
    /*
        是否橱窗推荐。 可选值：true，false。默认不过滤该条件     */
    HasShowcase  *bool `json:"has_showcase,omitempty" required:"false" `
    /*
        商品是否在淘宝显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" required:"false" `
    /*
        商品是否在外部网店显示     */
    IsEx  *bool `json:"is_ex,omitempty" required:"false" `
    /*
        起始的修改时间     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        结束的修改时间     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        是否挂接了达尔文标准产品体系     */
    IsCspu  *bool `json:"is_cspu,omitempty" required:"false" `
    /*
        组合商品     */
    IsCombine  *bool `json:"is_combine,omitempty" required:"false" `
    /*
        商品类型：a-拍卖,b-一口价     */
    AuctionType  *string `json:"auction_type,omitempty" required:"false" `
}

func (s *TaobaoItemsOnsaleGetRequest) SetFields(v []string) *TaobaoItemsOnsaleGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetQ(v string) *TaobaoItemsOnsaleGetRequest {
    s.Q = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetCid(v int64) *TaobaoItemsOnsaleGetRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetSellerCids(v string) *TaobaoItemsOnsaleGetRequest {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetPageNo(v int64) *TaobaoItemsOnsaleGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetPageSize(v int64) *TaobaoItemsOnsaleGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetOrderBy(v string) *TaobaoItemsOnsaleGetRequest {
    s.OrderBy = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetHasDiscount(v bool) *TaobaoItemsOnsaleGetRequest {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetHasShowcase(v bool) *TaobaoItemsOnsaleGetRequest {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetIsTaobao(v bool) *TaobaoItemsOnsaleGetRequest {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetIsEx(v bool) *TaobaoItemsOnsaleGetRequest {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetStartModified(v util.LocalTime) *TaobaoItemsOnsaleGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetEndModified(v util.LocalTime) *TaobaoItemsOnsaleGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetIsCspu(v bool) *TaobaoItemsOnsaleGetRequest {
    s.IsCspu = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetIsCombine(v bool) *TaobaoItemsOnsaleGetRequest {
    s.IsCombine = &v
    return s
}
func (s *TaobaoItemsOnsaleGetRequest) SetAuctionType(v string) *TaobaoItemsOnsaleGetRequest {
    s.AuctionType = &v
    return s
}

func (req *TaobaoItemsOnsaleGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Q != nil) {
        paramMap["q"] = *req.Q
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.SellerCids != nil) {
        paramMap["seller_cids"] = *req.SellerCids
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.OrderBy != nil) {
        paramMap["order_by"] = *req.OrderBy
    }
    if(req.HasDiscount != nil) {
        paramMap["has_discount"] = *req.HasDiscount
    }
    if(req.HasShowcase != nil) {
        paramMap["has_showcase"] = *req.HasShowcase
    }
    if(req.IsTaobao != nil) {
        paramMap["is_taobao"] = *req.IsTaobao
    }
    if(req.IsEx != nil) {
        paramMap["is_ex"] = *req.IsEx
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.IsCspu != nil) {
        paramMap["is_cspu"] = *req.IsCspu
    }
    if(req.IsCombine != nil) {
        paramMap["is_combine"] = *req.IsCombine
    }
    if(req.AuctionType != nil) {
        paramMap["auction_type"] = *req.AuctionType
    }
    return paramMap
}

func (req *TaobaoItemsOnsaleGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}