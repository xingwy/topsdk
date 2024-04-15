package domain


type TaobaoItemsSellerListGetCuntaoItemSpecific struct {
    /*
        村淘商品级佣金率     */
    KickBackRate  *string `json:"kick_back_rate,omitempty" `

}

func (s *TaobaoItemsSellerListGetCuntaoItemSpecific) SetKickBackRate(v string) *TaobaoItemsSellerListGetCuntaoItemSpecific {
    s.KickBackRate = &v
    return s
}
