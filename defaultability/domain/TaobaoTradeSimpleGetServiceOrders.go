package domain


type TaobaoTradeSimpleGetServiceOrders struct {
    /*
        虚拟服务子订单订单号     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        服务所属的交易订单号。如果服务为一年包换，则item_oid这笔订单享受改服务的保护     */
    ItemOid  *int64 `json:"item_oid,omitempty" `

    /*
        服务数字id     */
    ServiceId  *int64 `json:"service_id,omitempty" `

    /*
        服务详情的URL地址     */
    ServiceDetailUrl  *string `json:"service_detail_url,omitempty" `

    /*
        购买数量，取值范围为大于0的整数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        服务价格，精确到小数点后两位：单位:元     */
    Price  *string `json:"price,omitempty" `

    /*
        子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。     */
    Payment  *string `json:"payment,omitempty" `

    /*
        商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        服务子订单总费用     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        最近退款的id     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        服务图片地址     */
    PicPath  *string `json:"pic_path,omitempty" `

    /*
        支持家装类物流的类型     */
    TmserSpuCode  *string `json:"tmser_spu_code,omitempty" `

}

func (s *TaobaoTradeSimpleGetServiceOrders) SetOid(v int64) *TaobaoTradeSimpleGetServiceOrders {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetItemOid(v int64) *TaobaoTradeSimpleGetServiceOrders {
    s.ItemOid = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetServiceId(v int64) *TaobaoTradeSimpleGetServiceOrders {
    s.ServiceId = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetServiceDetailUrl(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.ServiceDetailUrl = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetNum(v int64) *TaobaoTradeSimpleGetServiceOrders {
    s.Num = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetPrice(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.Price = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetPayment(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetTitle(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.Title = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetTotalFee(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetRefundId(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetSellerNick(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetPicPath(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeSimpleGetServiceOrders) SetTmserSpuCode(v string) *TaobaoTradeSimpleGetServiceOrders {
    s.TmserSpuCode = &v
    return s
}
