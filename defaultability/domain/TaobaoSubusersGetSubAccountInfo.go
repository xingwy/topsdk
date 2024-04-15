package domain


type TaobaoSubusersGetSubAccountInfo struct {
    /*
        123456     */
    SubId  *int64 `json:"sub_id,omitempty" `

    /*
        1     */
    SubStatus  *int64 `json:"sub_status,omitempty" `

    /*
        654321     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        true     */
    SubDispatchStatus  *bool `json:"sub_dispatch_status,omitempty" `

    /*
        true     */
    SubOwedStatus  *bool `json:"sub_owed_status,omitempty" `

    /*
        zhangsan:no1     */
    SubNick  *string `json:"sub_nick,omitempty" `

    /*
        zhangsan     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        小红     */
    SubName  *string `json:"sub_name,omitempty" `

}

func (s *TaobaoSubusersGetSubAccountInfo) SetSubId(v int64) *TaobaoSubusersGetSubAccountInfo {
    s.SubId = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetSubStatus(v int64) *TaobaoSubusersGetSubAccountInfo {
    s.SubStatus = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetUserId(v int64) *TaobaoSubusersGetSubAccountInfo {
    s.UserId = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetSubDispatchStatus(v bool) *TaobaoSubusersGetSubAccountInfo {
    s.SubDispatchStatus = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetSubOwedStatus(v bool) *TaobaoSubusersGetSubAccountInfo {
    s.SubOwedStatus = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetSubNick(v string) *TaobaoSubusersGetSubAccountInfo {
    s.SubNick = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetUserNick(v string) *TaobaoSubusersGetSubAccountInfo {
    s.UserNick = &v
    return s
}
func (s *TaobaoSubusersGetSubAccountInfo) SetSubName(v string) *TaobaoSubusersGetSubAccountInfo {
    s.SubName = &v
    return s
}
