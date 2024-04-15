package domain

import (
        "topsdk/util"
    )

type TaobaoTmcUserGetTmcUser struct {
    /*
        用户首次开通时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        接收用户消息的组名     */
    GroupName  *string `json:"group_name,omitempty" `

    /*
        用户授权是否有效，true表示授权有效，false表示授权过期     */
    IsValid  *bool `json:"is_valid,omitempty" `

    /*
        用户最后开通时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        用户开通的消息类型列表。如果为空表示应用开通的所有类型     */
    Topics  *[]string `json:"topics,omitempty" `

    /*
        用户ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        用户昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户     */
    UserPlatform  *string `json:"user_platform,omitempty" `

}

func (s *TaobaoTmcUserGetTmcUser) SetCreated(v util.LocalTime) *TaobaoTmcUserGetTmcUser {
    s.Created = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetGroupName(v string) *TaobaoTmcUserGetTmcUser {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetIsValid(v bool) *TaobaoTmcUserGetTmcUser {
    s.IsValid = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetModified(v util.LocalTime) *TaobaoTmcUserGetTmcUser {
    s.Modified = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetTopics(v []string) *TaobaoTmcUserGetTmcUser {
    s.Topics = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetUserId(v int64) *TaobaoTmcUserGetTmcUser {
    s.UserId = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetUserNick(v string) *TaobaoTmcUserGetTmcUser {
    s.UserNick = &v
    return s
}
func (s *TaobaoTmcUserGetTmcUser) SetUserPlatform(v string) *TaobaoTmcUserGetTmcUser {
    s.UserPlatform = &v
    return s
}
