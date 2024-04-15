package domain


type TaobaoUserOpenuidGetbynickOpenUidInfo struct {
    /*
        买家openuid     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

    /*
        买家nick     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

}

func (s *TaobaoUserOpenuidGetbynickOpenUidInfo) SetBuyerOpenUid(v string) *TaobaoUserOpenuidGetbynickOpenUidInfo {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoUserOpenuidGetbynickOpenUidInfo) SetBuyerNick(v string) *TaobaoUserOpenuidGetbynickOpenUidInfo {
    s.BuyerNick = &v
    return s
}
