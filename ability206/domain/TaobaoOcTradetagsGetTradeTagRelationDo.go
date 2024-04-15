package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoOcTradetagsGetTradeTagRelationDo struct {
	/*
	   订单标签记录id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   订单id     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   标签名称     */
	TagName *string `json:"tag_name,omitempty" `

	/*
	   标签类型       1：官方标签      2：自定义标签     3：主站只读标签     */
	TagType *int64 `json:"tag_type,omitempty" `

	/*
	   标签值，json格式     */
	TagValue *string `json:"tag_value,omitempty" `

	/*
	   记录的创建时间     */
	GmtCreated *util.LocalTime `json:"gmt_created,omitempty" `

	/*
	   记录的最新修改时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   该标签在消费者端是否显示,0:不显示,1：显示     */
	Visible *int64 `json:"visible,omitempty" `

	/*
	   该标签操作的历史记录     */
	HistoryTradeTagRelations *[]TaobaoOcTradetagsGetHistoryTradeRelationDo `json:"history_trade_tag_relations,omitempty" `
}

func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetId(v int64) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.Id = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetTid(v int64) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.Tid = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetTagName(v string) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.TagName = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetTagType(v int64) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.TagType = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetTagValue(v string) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.TagValue = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetGmtCreated(v util.LocalTime) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.GmtCreated = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetGmtModified(v util.LocalTime) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.GmtModified = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetVisible(v int64) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.Visible = &v
	return s
}
func (s *TaobaoOcTradetagsGetTradeTagRelationDo) SetHistoryTradeTagRelations(v []TaobaoOcTradetagsGetHistoryTradeRelationDo) *TaobaoOcTradetagsGetTradeTagRelationDo {
	s.HistoryTradeTagRelations = &v
	return s
}
