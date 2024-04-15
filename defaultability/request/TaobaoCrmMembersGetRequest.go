package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmMembersGetRequest struct {
	/*
	   买家的昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" required:"false" `
	/*
	   会员等级,如果不传入值则默认为全部等级。 defalutValue��-1    */
	Grade *int64 `json:"grade,omitempty" required:"false" `
	/*
	   表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，     */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000     */
	CurrentPage *int64 `json:"current_page" required:"true" `
	/*
	   最迟上次交易时间     */
	MaxLastTradeTime *util.LocalTime `json:"max_last_trade_time,omitempty" required:"false" `
	/*
	   最早上次交易时间     */
	MinLastTradeTime *util.LocalTime `json:"min_last_trade_time,omitempty" required:"false" `
	/*
	   最大交易额，单位为元     */
	MaxTradeAmount *string `json:"max_trade_amount,omitempty" required:"false" `
	/*
	   最小交易额,单位为元     */
	MinTradeAmount *string `json:"min_trade_amount,omitempty" required:"false" `
	/*
	   最小交易量     */
	MinTradeCount *int64 `json:"min_trade_count,omitempty" required:"false" `
	/*
	   最大交易量     */
	MaxTradeCount *int64 `json:"max_trade_count,omitempty" required:"false" `
	/*
	   用户的ouid     */
	OpenUid *string `json:"open_uid,omitempty" required:"false" `
}

func (s *TaobaoCrmMembersGetRequest) SetBuyerNick(v string) *TaobaoCrmMembersGetRequest {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetGrade(v int64) *TaobaoCrmMembersGetRequest {
	s.Grade = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetPageSize(v int64) *TaobaoCrmMembersGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetCurrentPage(v int64) *TaobaoCrmMembersGetRequest {
	s.CurrentPage = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMaxLastTradeTime(v util.LocalTime) *TaobaoCrmMembersGetRequest {
	s.MaxLastTradeTime = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMinLastTradeTime(v util.LocalTime) *TaobaoCrmMembersGetRequest {
	s.MinLastTradeTime = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMaxTradeAmount(v string) *TaobaoCrmMembersGetRequest {
	s.MaxTradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMinTradeAmount(v string) *TaobaoCrmMembersGetRequest {
	s.MinTradeAmount = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMinTradeCount(v int64) *TaobaoCrmMembersGetRequest {
	s.MinTradeCount = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetMaxTradeCount(v int64) *TaobaoCrmMembersGetRequest {
	s.MaxTradeCount = &v
	return s
}
func (s *TaobaoCrmMembersGetRequest) SetOpenUid(v string) *TaobaoCrmMembersGetRequest {
	s.OpenUid = &v
	return s
}

func (req *TaobaoCrmMembersGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BuyerNick != nil {
		paramMap["buyer_nick"] = *req.BuyerNick
	}
	if req.Grade != nil {
		paramMap["grade"] = *req.Grade
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.CurrentPage != nil {
		paramMap["current_page"] = *req.CurrentPage
	}
	if req.MaxLastTradeTime != nil {
		paramMap["max_last_trade_time"] = *req.MaxLastTradeTime
	}
	if req.MinLastTradeTime != nil {
		paramMap["min_last_trade_time"] = *req.MinLastTradeTime
	}
	if req.MaxTradeAmount != nil {
		paramMap["max_trade_amount"] = *req.MaxTradeAmount
	}
	if req.MinTradeAmount != nil {
		paramMap["min_trade_amount"] = *req.MinTradeAmount
	}
	if req.MinTradeCount != nil {
		paramMap["min_trade_count"] = *req.MinTradeCount
	}
	if req.MaxTradeCount != nil {
		paramMap["max_trade_count"] = *req.MaxTradeCount
	}
	if req.OpenUid != nil {
		paramMap["open_uid"] = *req.OpenUid
	}
	return paramMap
}

func (req *TaobaoCrmMembersGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
