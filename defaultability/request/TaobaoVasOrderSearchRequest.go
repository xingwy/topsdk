package request

import (
        "topsdk/util"
    )

type TaobaoVasOrderSearchRequest struct {
    /*
        淘宝会员名     */
    Nick  *string `json:"nick,omitempty" required:"false" `
    /*
        应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码     */
    ArticleCode  *string `json:"article_code" required:"true" `
    /*
        收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码     */
    ItemCode  *string `json:"item_code,omitempty" required:"false" `
    /*
        订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）     */
    StartCreated  *util.LocalTime `json:"start_created,omitempty" required:"false" `
    /*
        订单创建时间（订购时间）结束值     */
    EndCreated  *util.LocalTime `json:"end_created,omitempty" required:"false" `
    /*
        订单号 defalutValue��0    */
    BizOrderId  *int64 `json:"biz_order_id,omitempty" required:"false" `
    /*
        子订单号 defalutValue��0    */
    OrderId  *int64 `json:"order_id,omitempty" required:"false" `
    /*
        订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部 defalutValue��0    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        一页包含的记录数     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoVasOrderSearchRequest) SetNick(v string) *TaobaoVasOrderSearchRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetArticleCode(v string) *TaobaoVasOrderSearchRequest {
    s.ArticleCode = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetItemCode(v string) *TaobaoVasOrderSearchRequest {
    s.ItemCode = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetStartCreated(v util.LocalTime) *TaobaoVasOrderSearchRequest {
    s.StartCreated = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetEndCreated(v util.LocalTime) *TaobaoVasOrderSearchRequest {
    s.EndCreated = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetBizOrderId(v int64) *TaobaoVasOrderSearchRequest {
    s.BizOrderId = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetOrderId(v int64) *TaobaoVasOrderSearchRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetBizType(v int64) *TaobaoVasOrderSearchRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetPageSize(v int64) *TaobaoVasOrderSearchRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoVasOrderSearchRequest) SetPageNo(v int64) *TaobaoVasOrderSearchRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoVasOrderSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.ArticleCode != nil) {
        paramMap["article_code"] = *req.ArticleCode
    }
    if(req.ItemCode != nil) {
        paramMap["item_code"] = *req.ItemCode
    }
    if(req.StartCreated != nil) {
        paramMap["start_created"] = *req.StartCreated
    }
    if(req.EndCreated != nil) {
        paramMap["end_created"] = *req.EndCreated
    }
    if(req.BizOrderId != nil) {
        paramMap["biz_order_id"] = *req.BizOrderId
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoVasOrderSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}