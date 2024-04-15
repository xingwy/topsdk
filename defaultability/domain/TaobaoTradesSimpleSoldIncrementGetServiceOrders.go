package domain


type TaobaoTradesSimpleSoldIncrementGetServiceOrders struct {
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
    RefundId  *int64 `json:"refund_id,omitempty" `

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

func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetOid(v int64) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.Oid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetItemOid(v int64) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.ItemOid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetServiceId(v int64) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.ServiceId = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetNum(v int64) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetPrice(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetPayment(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetTitle(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetTotalFee(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetRefundId(v int64) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetSellerNick(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetPicPath(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSimpleSoldIncrementGetServiceOrders) SetTmserSpuCode(v string) *TaobaoTradesSimpleSoldIncrementGetServiceOrders {
    s.TmserSpuCode = &v
    return s
}
