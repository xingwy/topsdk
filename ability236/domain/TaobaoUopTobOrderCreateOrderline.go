package domain


type TaobaoUopTobOrderCreateOrderline struct {
    /*
        订单行号     */
    OrderLineNo  *string `json:"order_line_no,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        商品编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        数量     */
    Amount  *string `json:"amount,omitempty" `

}

func (s *TaobaoUopTobOrderCreateOrderline) SetOrderLineNo(v string) *TaobaoUopTobOrderCreateOrderline {
    s.OrderLineNo = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderline) SetItemId(v string) *TaobaoUopTobOrderCreateOrderline {
    s.ItemId = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderline) SetItemCode(v string) *TaobaoUopTobOrderCreateOrderline {
    s.ItemCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderline) SetAmount(v string) *TaobaoUopTobOrderCreateOrderline {
    s.Amount = &v
    return s
}
