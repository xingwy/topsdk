package domain


type TaobaoUopTobOrderCreateDeliveryorder struct {
    /*
        物流单号     */
    CnOrderCode  *string `json:"cn_order_code,omitempty" `

    /*
        仓库编码     */
    WarehouseCode  *string `json:"warehouse_code,omitempty" `

    /*
        创建时间     */
    CreateTime  *string `json:"create_time,omitempty" `

    /*
        订单行     */
    OrderLines  *[]TaobaoUopTobOrderCreateOrderline `json:"order_lines,omitempty" `

}

func (s *TaobaoUopTobOrderCreateDeliveryorder) SetCnOrderCode(v string) *TaobaoUopTobOrderCreateDeliveryorder {
    s.CnOrderCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryorder) SetWarehouseCode(v string) *TaobaoUopTobOrderCreateDeliveryorder {
    s.WarehouseCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryorder) SetCreateTime(v string) *TaobaoUopTobOrderCreateDeliveryorder {
    s.CreateTime = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryorder) SetOrderLines(v []TaobaoUopTobOrderCreateOrderline) *TaobaoUopTobOrderCreateDeliveryorder {
    s.OrderLines = &v
    return s
}
