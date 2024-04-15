package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoJushitaJmsUserGetTmcUser struct {
	/*
	   用户首次开通时间     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   用户授权是否有效，true表示授权有效，false表示授权过期     */
	IsValid *bool `json:"is_valid,omitempty" `

	/*
	   用户最后开通时间     */
	Modified *util.LocalTime `json:"modified,omitempty" `

	/*
	   用户ID     */
	UserId *int64 `json:"user_id,omitempty" `

	/*
	   用户昵称     */
	UserNick *string `json:"user_nick,omitempty" `
}

func (s *TaobaoJushitaJmsUserGetTmcUser) SetCreated(v util.LocalTime) *TaobaoJushitaJmsUserGetTmcUser {
	s.Created = &v
	return s
}
func (s *TaobaoJushitaJmsUserGetTmcUser) SetIsValid(v bool) *TaobaoJushitaJmsUserGetTmcUser {
	s.IsValid = &v
	return s
}
func (s *TaobaoJushitaJmsUserGetTmcUser) SetModified(v util.LocalTime) *TaobaoJushitaJmsUserGetTmcUser {
	s.Modified = &v
	return s
}
func (s *TaobaoJushitaJmsUserGetTmcUser) SetUserId(v int64) *TaobaoJushitaJmsUserGetTmcUser {
	s.UserId = &v
	return s
}
func (s *TaobaoJushitaJmsUserGetTmcUser) SetUserNick(v string) *TaobaoJushitaJmsUserGetTmcUser {
	s.UserNick = &v
	return s
}
