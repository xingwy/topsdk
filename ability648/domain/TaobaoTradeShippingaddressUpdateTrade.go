package domain

import (
        "topsdk/util"
    )

type TaobaoTradeShippingaddressUpdateTrade struct {
    /*
        交易编号 (父订单的交易编号)     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoTradeShippingaddressUpdateTrade) SetTid(v int64) *TaobaoTradeShippingaddressUpdateTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateTrade) SetModified(v util.LocalTime) *TaobaoTradeShippingaddressUpdateTrade {
    s.Modified = &v
    return s
}
