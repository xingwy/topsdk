package request

import (
        "topsdk/util"
    )

type TaobaoItemsInventoryGetRequest struct {
    /*
        需返回的字段列表。可选值为 Item 商品结构体中的以下字段： 
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。     */
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
        页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数。取值范围:大于零的整数;最大值：200；默认值：40。 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc defalutValue��list_time:desc    */
    OrderBy  *string `json:"order_by,omitempty" required:"false" `
    /*
        是否参与会员折扣。可选值：true，false。默认不过滤该条件     */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
    /*
        商品是否在淘宝显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" required:"false" `
    /*
        商品是否在外部网店显示     */
    IsEx  *bool `json:"is_ex,omitempty" required:"false" `
    /*
        分类字段。可选值:<br>
regular_shelved(定时上架)<br>
never_on_shelf(从未上架)<br>
off_shelf(我下架的)<br>
<font color='red'>for_shelved(等待所有上架)<br>
sold_out(全部卖完)<br>
violation_off_shelf(违规下架的)<br>
默认查询for_shelved(等待所有上架)这个状态的商品<br></font>
注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)     */
    Banner  *string `json:"banner,omitempty" required:"false" `
    /*
        商品起始修改时间     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        商品结束修改时间     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        商品类型：a-拍卖,b-一口价     */
    AuctionType  *string `json:"auction_type,omitempty" required:"false" `
}

func (s *TaobaoItemsInventoryGetRequest) SetFields(v []string) *TaobaoItemsInventoryGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetQ(v string) *TaobaoItemsInventoryGetRequest {
    s.Q = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetCid(v int64) *TaobaoItemsInventoryGetRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetSellerCids(v string) *TaobaoItemsInventoryGetRequest {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetPageNo(v int64) *TaobaoItemsInventoryGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetPageSize(v int64) *TaobaoItemsInventoryGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetOrderBy(v string) *TaobaoItemsInventoryGetRequest {
    s.OrderBy = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetHasDiscount(v bool) *TaobaoItemsInventoryGetRequest {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetIsTaobao(v bool) *TaobaoItemsInventoryGetRequest {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetIsEx(v bool) *TaobaoItemsInventoryGetRequest {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetBanner(v string) *TaobaoItemsInventoryGetRequest {
    s.Banner = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetStartModified(v util.LocalTime) *TaobaoItemsInventoryGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetEndModified(v util.LocalTime) *TaobaoItemsInventoryGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoItemsInventoryGetRequest) SetAuctionType(v string) *TaobaoItemsInventoryGetRequest {
    s.AuctionType = &v
    return s
}

func (req *TaobaoItemsInventoryGetRequest) ToMap() map[string]interface{} {
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
    if(req.IsTaobao != nil) {
        paramMap["is_taobao"] = *req.IsTaobao
    }
    if(req.IsEx != nil) {
        paramMap["is_ex"] = *req.IsEx
    }
    if(req.Banner != nil) {
        paramMap["banner"] = *req.Banner
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.AuctionType != nil) {
        paramMap["auction_type"] = *req.AuctionType
    }
    return paramMap
}

func (req *TaobaoItemsInventoryGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}