package domain

import (
        "topsdk/util"
    )

type TaobaoJdsTradeTracesGetTradeTrace struct {
    /*
        动作发生的时间     */
    ActionTime  *util.LocalTime `json:"action_time,omitempty" `

    /*
        应用标题     */
    AppTitle  *string `json:"app_title,omitempty" `

    /*
        子订单的id列表,以逗号分割     */
    OrderIds  *string `json:"order_ids,omitempty" `

    /*
        备注字段     */
    Remark  *string `json:"remark,omitempty" `

    /*
        卖家的淘宝nick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        回流的订单状态     */
    Status  *string `json:"status,omitempty" `

    /*
        交易tid     */
    Tid  *int64 `json:"tid,omitempty" `

}

func (s *TaobaoJdsTradeTracesGetTradeTrace) SetActionTime(v util.LocalTime) *TaobaoJdsTradeTracesGetTradeTrace {
    s.ActionTime = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetAppTitle(v string) *TaobaoJdsTradeTracesGetTradeTrace {
    s.AppTitle = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetOrderIds(v string) *TaobaoJdsTradeTracesGetTradeTrace {
    s.OrderIds = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetRemark(v string) *TaobaoJdsTradeTracesGetTradeTrace {
    s.Remark = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetSellerNick(v string) *TaobaoJdsTradeTracesGetTradeTrace {
    s.SellerNick = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetStatus(v string) *TaobaoJdsTradeTracesGetTradeTrace {
    s.Status = &v
    return s
}
func (s *TaobaoJdsTradeTracesGetTradeTrace) SetTid(v int64) *TaobaoJdsTradeTracesGetTradeTrace {
    s.Tid = &v
    return s
}
