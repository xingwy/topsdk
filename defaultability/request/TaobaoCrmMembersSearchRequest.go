package request

import (
        "topsdk/util"
    )

type TaobaoCrmMembersSearchRequest struct {
    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" required:"false" `
    /*
        会员等级 defalutValue��-1    */
    Grade  *int64 `json:"grade,omitempty" required:"false" `
    /*
        关系来源，1交易成功，2未成交，3卖家手动吸纳     */
    RelationSource  *int64 `json:"relation_source,omitempty" required:"false" `
    /*
        最迟上次交易时间     */
    MaxLastTradeTime  *util.LocalTime `json:"max_last_trade_time,omitempty" required:"false" `
    /*
        最早上次交易时间（订单创建时间）     */
    MinLastTradeTime  *util.LocalTime `json:"min_last_trade_time,omitempty" required:"false" `
    /*
        最小交易量     */
    MinTradeCount  *int64 `json:"min_trade_count,omitempty" required:"false" `
    /*
        最大交易量     */
    MaxTradeCount  *int64 `json:"max_trade_count,omitempty" required:"false" `
    /*
        分组id     */
    GroupId  *int64 `json:"group_id,omitempty" required:"false" `
    /*
        显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大40页     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
    /*
        每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        最小交易额，单位为元     */
    MinTradeAmount  *string `json:"min_trade_amount,omitempty" required:"false" `
    /*
        最大交易额，单位为元     */
    MaxTradeAmount  *string `json:"max_trade_amount,omitempty" required:"false" `
    /*
        用户的open_uid     */
    OpenUid  *string `json:"open_uid,omitempty" required:"false" `
}

func (s *TaobaoCrmMembersSearchRequest) SetBuyerNick(v string) *TaobaoCrmMembersSearchRequest {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetGrade(v int64) *TaobaoCrmMembersSearchRequest {
    s.Grade = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetRelationSource(v int64) *TaobaoCrmMembersSearchRequest {
    s.RelationSource = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMaxLastTradeTime(v util.LocalTime) *TaobaoCrmMembersSearchRequest {
    s.MaxLastTradeTime = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMinLastTradeTime(v util.LocalTime) *TaobaoCrmMembersSearchRequest {
    s.MinLastTradeTime = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMinTradeCount(v int64) *TaobaoCrmMembersSearchRequest {
    s.MinTradeCount = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMaxTradeCount(v int64) *TaobaoCrmMembersSearchRequest {
    s.MaxTradeCount = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetGroupId(v int64) *TaobaoCrmMembersSearchRequest {
    s.GroupId = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetCurrentPage(v int64) *TaobaoCrmMembersSearchRequest {
    s.CurrentPage = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetPageSize(v int64) *TaobaoCrmMembersSearchRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMinTradeAmount(v string) *TaobaoCrmMembersSearchRequest {
    s.MinTradeAmount = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetMaxTradeAmount(v string) *TaobaoCrmMembersSearchRequest {
    s.MaxTradeAmount = &v
    return s
}
func (s *TaobaoCrmMembersSearchRequest) SetOpenUid(v string) *TaobaoCrmMembersSearchRequest {
    s.OpenUid = &v
    return s
}

func (req *TaobaoCrmMembersSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    if(req.Grade != nil) {
        paramMap["grade"] = *req.Grade
    }
    if(req.RelationSource != nil) {
        paramMap["relation_source"] = *req.RelationSource
    }
    if(req.MaxLastTradeTime != nil) {
        paramMap["max_last_trade_time"] = *req.MaxLastTradeTime
    }
    if(req.MinLastTradeTime != nil) {
        paramMap["min_last_trade_time"] = *req.MinLastTradeTime
    }
    if(req.MinTradeCount != nil) {
        paramMap["min_trade_count"] = *req.MinTradeCount
    }
    if(req.MaxTradeCount != nil) {
        paramMap["max_trade_count"] = *req.MaxTradeCount
    }
    if(req.GroupId != nil) {
        paramMap["group_id"] = *req.GroupId
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.MinTradeAmount != nil) {
        paramMap["min_trade_amount"] = *req.MinTradeAmount
    }
    if(req.MaxTradeAmount != nil) {
        paramMap["max_trade_amount"] = *req.MaxTradeAmount
    }
    if(req.OpenUid != nil) {
        paramMap["open_uid"] = *req.OpenUid
    }
    return paramMap
}

func (req *TaobaoCrmMembersSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}