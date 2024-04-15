package domain


type AlibabaAscpLogisticsCpGetLogisticsResourceRequest struct {
    /*
        可选值:offline(自己联系发货),online(在线下单),all(自己联系+在线下单)instant(同城配送).  defalutValue:offline    */
    OrderMode  *string `json:"order_mode,omitempty" `

}

func (s *AlibabaAscpLogisticsCpGetLogisticsResourceRequest) SetOrderMode(v string) *AlibabaAscpLogisticsCpGetLogisticsResourceRequest {
    s.OrderMode = &v
    return s
}
