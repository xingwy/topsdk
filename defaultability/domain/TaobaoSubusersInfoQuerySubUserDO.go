package domain


type TaobaoSubusersInfoQuerySubUserDO struct {
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

}

func (s *TaobaoSubusersInfoQuerySubUserDO) SetSubId(v int64) *TaobaoSubusersInfoQuerySubUserDO {
    s.SubId = &v
    return s
}
func (s *TaobaoSubusersInfoQuerySubUserDO) SetStatus(v int64) *TaobaoSubusersInfoQuerySubUserDO {
    s.Status = &v
    return s
}
func (s *TaobaoSubusersInfoQuerySubUserDO) SetSellerId(v int64) *TaobaoSubusersInfoQuerySubUserDO {
    s.SellerId = &v
    return s
}
func (s *TaobaoSubusersInfoQuerySubUserDO) SetNick(v string) *TaobaoSubusersInfoQuerySubUserDO {
    s.Nick = &v
    return s
}
func (s *TaobaoSubusersInfoQuerySubUserDO) SetSellerNick(v string) *TaobaoSubusersInfoQuerySubUserDO {
    s.SellerNick = &v
    return s
}
