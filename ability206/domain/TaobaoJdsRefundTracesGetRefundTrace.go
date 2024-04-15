package domain

import (
        "topsdk/util"
    )

type TaobaoJdsRefundTracesGetRefundTrace struct {
    /*
        动作发生的时间     */
    ActionTime  *util.LocalTime `json:"action_time,omitempty" `

    /*
        应用标题     */
    AppTitle  *string `json:"app_title,omitempty" `

    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id,omitempty" `

    /*
        备注字段     */
    Remark  *string `json:"remark,omitempty" `

    /*
        回流的订单状态     */
    Status  *string `json:"status,omitempty" `

    /*
        交易tid     */
    Tid  *int64 `json:"tid,omitempty" `

}

func (s *TaobaoJdsRefundTracesGetRefundTrace) SetActionTime(v util.LocalTime) *TaobaoJdsRefundTracesGetRefundTrace {
    s.ActionTime = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRefundTrace) SetAppTitle(v string) *TaobaoJdsRefundTracesGetRefundTrace {
    s.AppTitle = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRefundTrace) SetRefundId(v int64) *TaobaoJdsRefundTracesGetRefundTrace {
    s.RefundId = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRefundTrace) SetRemark(v string) *TaobaoJdsRefundTracesGetRefundTrace {
    s.Remark = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRefundTrace) SetStatus(v string) *TaobaoJdsRefundTracesGetRefundTrace {
    s.Status = &v
    return s
}
func (s *TaobaoJdsRefundTracesGetRefundTrace) SetTid(v int64) *TaobaoJdsRefundTracesGetRefundTrace {
    s.Tid = &v
    return s
}
