package domain


type TaobaoItemSellerGetCuntaoItemSpecific struct {
    /*
        村淘商品级佣金率     */
    KickBackRate  *string `json:"kick_back_rate,omitempty" `

}

func (s *TaobaoItemSellerGetCuntaoItemSpecific) SetKickBackRate(v string) *TaobaoItemSellerGetCuntaoItemSpecific {
    s.KickBackRate = &v
    return s
}
