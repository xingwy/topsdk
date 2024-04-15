package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoVasSubscSearchRequest struct {
	/*
	   淘宝会员名     */
	Nick *string `json:"nick,omitempty" required:"false" `
	/*
	   应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码     */
	ArticleCode *string `json:"article_code" required:"true" `
	/*
	   收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码     */
	ItemCode *string `json:"item_code,omitempty" required:"false" `
	/*
	   到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）     */
	StartDeadline *util.LocalTime `json:"start_deadline,omitempty" required:"false" `
	/*
	   到期时间结束值     */
	EndDeadline *util.LocalTime `json:"end_deadline,omitempty" required:"false" `
	/*
	   订购记录状态，1=有效 2=过期 空=全部     */
	Status *int64 `json:"status,omitempty" required:"false" `
	/*
	   是否自动续费，true=自动续费 false=非自动续费 空=全部     */
	Autosub *bool `json:"autosub,omitempty" required:"false" `
	/*
	   是否到期提醒，true=到期提醒 false=非到期提醒 空=全部     */
	ExpireNotice *bool `json:"expire_notice,omitempty" required:"false" `
	/*
	   一页包含的记录数     */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   页码     */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoVasSubscSearchRequest) SetNick(v string) *TaobaoVasSubscSearchRequest {
	s.Nick = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetArticleCode(v string) *TaobaoVasSubscSearchRequest {
	s.ArticleCode = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetItemCode(v string) *TaobaoVasSubscSearchRequest {
	s.ItemCode = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetStartDeadline(v util.LocalTime) *TaobaoVasSubscSearchRequest {
	s.StartDeadline = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetEndDeadline(v util.LocalTime) *TaobaoVasSubscSearchRequest {
	s.EndDeadline = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetStatus(v int64) *TaobaoVasSubscSearchRequest {
	s.Status = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetAutosub(v bool) *TaobaoVasSubscSearchRequest {
	s.Autosub = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetExpireNotice(v bool) *TaobaoVasSubscSearchRequest {
	s.ExpireNotice = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetPageSize(v int64) *TaobaoVasSubscSearchRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoVasSubscSearchRequest) SetPageNo(v int64) *TaobaoVasSubscSearchRequest {
	s.PageNo = &v
	return s
}

func (req *TaobaoVasSubscSearchRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Nick != nil {
		paramMap["nick"] = *req.Nick
	}
	if req.ArticleCode != nil {
		paramMap["article_code"] = *req.ArticleCode
	}
	if req.ItemCode != nil {
		paramMap["item_code"] = *req.ItemCode
	}
	if req.StartDeadline != nil {
		paramMap["start_deadline"] = *req.StartDeadline
	}
	if req.EndDeadline != nil {
		paramMap["end_deadline"] = *req.EndDeadline
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.Autosub != nil {
		paramMap["autosub"] = *req.Autosub
	}
	if req.ExpireNotice != nil {
		paramMap["expire_notice"] = *req.ExpireNotice
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	return paramMap
}

func (req *TaobaoVasSubscSearchRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
