package domain


type TaobaoSubusersSubaccountSearchSubAccountInfo struct {
    /*
        子账号id     */
    SubId  *int64 `json:"sub_id,omitempty" `

    /*
        子账号状态， 1正常   2冻结  3处罚     */
    Status  *int64 `json:"status,omitempty" `

    /*
        主账号id     */
    SellerId  *string `json:"seller_id,omitempty" `

    /*
        子账号登录名     */
    SubNick  *string `json:"sub_nick,omitempty" `

}

func (s *TaobaoSubusersSubaccountSearchSubAccountInfo) SetSubId(v int64) *TaobaoSubusersSubaccountSearchSubAccountInfo {
    s.SubId = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchSubAccountInfo) SetStatus(v int64) *TaobaoSubusersSubaccountSearchSubAccountInfo {
    s.Status = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchSubAccountInfo) SetSellerId(v string) *TaobaoSubusersSubaccountSearchSubAccountInfo {
    s.SellerId = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchSubAccountInfo) SetSubNick(v string) *TaobaoSubusersSubaccountSearchSubAccountInfo {
    s.SubNick = &v
    return s
}
