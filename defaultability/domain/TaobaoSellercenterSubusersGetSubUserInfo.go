package domain


type TaobaoSellercenterSubusersGetSubUserInfo struct {
    /*
        是否参与分流 1不参与 2参与     */
    IsOnline  *int64 `json:"is_online,omitempty" `

    /*
        子账号Id     */
    SubId  *int64 `json:"sub_id,omitempty" `

    /*
        子账号当前状态 1正常 -1删除  2冻结     */
    Status  *int64 `json:"status,omitempty" `

    /*
        子账号所属的主账号的唯一标识     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        子账号用户名     */
    Nick  *string `json:"nick,omitempty" `

    /*
        主账号昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        子账号姓名     */
    FullName  *string `json:"full_name,omitempty" `

}

func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetIsOnline(v int64) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.IsOnline = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetSubId(v int64) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.SubId = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetStatus(v int64) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.Status = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetSellerId(v int64) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.SellerId = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetNick(v string) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.Nick = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetSellerNick(v string) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.SellerNick = &v
    return s
}
func (s *TaobaoSellercenterSubusersGetSubUserInfo) SetFullName(v string) *TaobaoSellercenterSubusersGetSubUserInfo {
    s.FullName = &v
    return s
}
