package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoVasSubscSearchArticleSub struct {
	/*
	   淘宝会员名     */
	Nick *string `json:"nick,omitempty" `

	/*
	   应用名称     */
	ArticleName *string `json:"article_name,omitempty" `

	/*
	   应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码     */
	ArticleCode *string `json:"article_code,omitempty" `

	/*
	   收费项目名称     */
	ItemName *string `json:"item_name,omitempty" `

	/*
	   收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码     */
	ItemCode *string `json:"item_code,omitempty" `

	/*
	   订购关系到期时间     */
	Deadline *util.LocalTime `json:"deadline,omitempty" `

	/*
	   状态，1=有效 2=过期     */
	Status *int64 `json:"status,omitempty" `

	/*
	   是否自动续费     */
	Autosub *bool `json:"autosub,omitempty" `

	/*
	   是否到期提醒     */
	ExpireNotice *bool `json:"expire_notice,omitempty" `

	/*
	   订购关系开始时间     */
	StartDate *util.LocalTime `json:"start_date,omitempty" `
}

func (s *TaobaoVasSubscSearchArticleSub) SetNick(v string) *TaobaoVasSubscSearchArticleSub {
	s.Nick = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetArticleName(v string) *TaobaoVasSubscSearchArticleSub {
	s.ArticleName = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetArticleCode(v string) *TaobaoVasSubscSearchArticleSub {
	s.ArticleCode = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetItemName(v string) *TaobaoVasSubscSearchArticleSub {
	s.ItemName = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetItemCode(v string) *TaobaoVasSubscSearchArticleSub {
	s.ItemCode = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetDeadline(v util.LocalTime) *TaobaoVasSubscSearchArticleSub {
	s.Deadline = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetStatus(v int64) *TaobaoVasSubscSearchArticleSub {
	s.Status = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetAutosub(v bool) *TaobaoVasSubscSearchArticleSub {
	s.Autosub = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetExpireNotice(v bool) *TaobaoVasSubscSearchArticleSub {
	s.ExpireNotice = &v
	return s
}
func (s *TaobaoVasSubscSearchArticleSub) SetStartDate(v util.LocalTime) *TaobaoVasSubscSearchArticleSub {
	s.StartDate = &v
	return s
}
