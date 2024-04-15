package domain


type TaobaoTradeSellermemosGetOrderSellerMemo struct {
    /*
        订单号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        订单备注     */
    Memo  *string `json:"memo,omitempty" `

    /*
        插旗ID     */
    FlagId  *int64 `json:"flag_id,omitempty" `

    /*
        插旗备注     */
    FlagTag  *string `json:"flag_tag,omitempty" `

    /*
        备注创建者ID     */
    OperatorId  *int64 `json:"operator_id,omitempty" `

    /*
        备注创建时间戳     */
    Timestamp  *int64 `json:"timestamp,omitempty" `

}

func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetTid(v int64) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetMemo(v string) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.Memo = &v
    return s
}
func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetFlagId(v int64) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.FlagId = &v
    return s
}
func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetFlagTag(v string) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.FlagTag = &v
    return s
}
func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetOperatorId(v int64) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.OperatorId = &v
    return s
}
func (s *TaobaoTradeSellermemosGetOrderSellerMemo) SetTimestamp(v int64) *TaobaoTradeSellermemosGetOrderSellerMemo {
    s.Timestamp = &v
    return s
}
