package domain


type TaobaoSubusersPageSubAccountInfo struct {
    /*
        子账号id     */
    SubId  *int64 `json:"sub_id,omitempty" `

    /*
        子账号状态， 1正常   2冻结  3处罚     */
    SubStatus  *int64 `json:"sub_status,omitempty" `

    /*
        主账号id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        1     */
    SubDispatchStatus  *bool `json:"sub_dispatch_status,omitempty" `

    /*
        1     */
    SubOwedStatus  *bool `json:"sub_owed_status,omitempty" `

    /*
        zhangsan:no1     */
    SubNick  *string `json:"sub_nick,omitempty" `

    /*
        zhangsan     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        员工昵称     */
    SubName  *string `json:"sub_name,omitempty" `

}

func (s *TaobaoSubusersPageSubAccountInfo) SetSubId(v int64) *TaobaoSubusersPageSubAccountInfo {
    s.SubId = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetSubStatus(v int64) *TaobaoSubusersPageSubAccountInfo {
    s.SubStatus = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetUserId(v int64) *TaobaoSubusersPageSubAccountInfo {
    s.UserId = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetSubDispatchStatus(v bool) *TaobaoSubusersPageSubAccountInfo {
    s.SubDispatchStatus = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetSubOwedStatus(v bool) *TaobaoSubusersPageSubAccountInfo {
    s.SubOwedStatus = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetSubNick(v string) *TaobaoSubusersPageSubAccountInfo {
    s.SubNick = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetUserNick(v string) *TaobaoSubusersPageSubAccountInfo {
    s.UserNick = &v
    return s
}
func (s *TaobaoSubusersPageSubAccountInfo) SetSubName(v string) *TaobaoSubusersPageSubAccountInfo {
    s.SubName = &v
    return s
}
