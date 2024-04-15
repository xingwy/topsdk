package domain


type TaobaoRefundStatusGetRefundQueryByOrderIdRequest struct {
    /*
        订单号     */
    BizOrderId  *int64 `json:"biz_order_id,omitempty" `

}

func (s *TaobaoRefundStatusGetRefundQueryByOrderIdRequest) SetBizOrderId(v int64) *TaobaoRefundStatusGetRefundQueryByOrderIdRequest {
    s.BizOrderId = &v
    return s
}
