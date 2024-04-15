package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TmallTraderateFeedsGetModel struct {
	/*
	   追加评价内容     */
	AppendContent *string `json:"append_content,omitempty" `

	/*
	   原始评价对应的标签列表     */
	Tags *[]TmallTraderateFeedsGetTags `json:"tags,omitempty" `

	/*
	   评价内容     */
	Content *string `json:"content,omitempty" `

	/*
	   追加评价中带有的语义标签列表     */
	AppendTags *[]TmallTraderateFeedsGetTags `json:"append_tags,omitempty" `

	/*
	   追加评价时间     */
	AppendTime *util.LocalTime `json:"append_time,omitempty" `

	/*
	   追评中是否含有负向标签     */
	AppendHasNegtv *bool `json:"append_has_negtv,omitempty" `

	/*
	   表示评价者的昵称     */
	UserNick *string `json:"user_nick,omitempty" `

	/*
	   原始评价是否含有负向标签     */
	HasNegtv *bool `json:"has_negtv,omitempty" `

	/*
	   评价时间     */
	CommentTime *util.LocalTime `json:"comment_time,omitempty" `

	/*
	   ouid     */
	Ouid *string `json:"ouid,omitempty" `
}

func (s *TmallTraderateFeedsGetModel) SetAppendContent(v string) *TmallTraderateFeedsGetModel {
	s.AppendContent = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetTags(v []TmallTraderateFeedsGetTags) *TmallTraderateFeedsGetModel {
	s.Tags = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetContent(v string) *TmallTraderateFeedsGetModel {
	s.Content = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetAppendTags(v []TmallTraderateFeedsGetTags) *TmallTraderateFeedsGetModel {
	s.AppendTags = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetAppendTime(v util.LocalTime) *TmallTraderateFeedsGetModel {
	s.AppendTime = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetAppendHasNegtv(v bool) *TmallTraderateFeedsGetModel {
	s.AppendHasNegtv = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetUserNick(v string) *TmallTraderateFeedsGetModel {
	s.UserNick = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetHasNegtv(v bool) *TmallTraderateFeedsGetModel {
	s.HasNegtv = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetCommentTime(v util.LocalTime) *TmallTraderateFeedsGetModel {
	s.CommentTime = &v
	return s
}
func (s *TmallTraderateFeedsGetModel) SetOuid(v string) *TmallTraderateFeedsGetModel {
	s.Ouid = &v
	return s
}
