package domain

import (
        "topsdk/util"
    )

type TaobaoTradeSimpleGetTrade struct {
    /*
        买家OpenUid     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

    /*
        交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据 可选值 fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) auto_delivery(自动发货) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易)o2o_offlinetrade（O2O交易）step (万人团)nopaid(无付款订单)pre_auth_type(预授权0元购机交易)     */
    Type  *string `json:"type,omitempty" `

    /*
        交易创建时间。格式:yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        交易编号 (父订单的交易编号)     */
    Tid  *string `json:"tid,omitempty" `

    /*
        卖家是否已评价。可选值:true(已评价),false(未评价)     */
    SellerRate  *bool `json:"seller_rate,omitempty" `

    /*
        买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false     */
    BuyerRate  *bool `json:"buyer_rate,omitempty" `

    /*
        交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)	* PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)     */
    Status  *string `json:"status,omitempty" `

    /*
        实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Payment  *string `json:"payment,omitempty" `

    /*
        建议使用trade.promotion_details查询系统优惠系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分     */
    DiscountFee  *string `json:"discount_fee,omitempty" `

    /*
        卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样     */
    AdjustFee  *string `json:"adjust_fee,omitempty" `

    /*
        邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    PostFee  *string `json:"post_fee,omitempty" `

    /*
        商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        卖家发货时间。格式:yyyy-MM-dd HH:mm:ss     */
    ConsignTime  *util.LocalTime `json:"consign_time,omitempty" `

    /*
        卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    ReceivedPayment  *string `json:"received_payment,omitempty" `

    /*
        买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段）     */
    BuyerMemo  *string `json:"buyer_memo,omitempty" `

    /*
        卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段）     */
    SellerMemo  *string `json:"seller_memo,omitempty" `

    /*
        买家留言     */
    BuyerMessage  *string `json:"buyer_message,omitempty" `

    /*
        商品图片绝对途径     */
    PicPath  *string `json:"pic_path,omitempty" `

    /*
        商品数字编号     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分     */
    Price  *string `json:"price,omitempty" `

    /*
        创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送)。     */
    ShippingType  *string `json:"shipping_type,omitempty" `

    /*
        商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。     */
    Num  *int64 `json:"num,omitempty" `

    /*
        收货人的所在省份     */
    ReceiverState  *string `json:"receiver_state,omitempty" `

    /*
        收货人的所在城市<br/>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br/>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        交易中剩余的确认收货金额（这个金额会随着子订单确认收货而不断减少，交易成功后会变为零）。精确到2位小数;单位:元。如:200.07，表示:200 元7分     */
    AvailableConfirmFee  *string `json:"available_confirm_fee,omitempty" `

    /*
        是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含)     */
    HasPostFee  *bool `json:"has_post_fee,omitempty" `

    /*
        超时到期时间。格式:yyyy-MM-dd HH:mm:ss。业务规则：前提条件：只有在买家已付款，卖家已发货的情况下才有效如果申请了退款，那么超时会落在子订单上；比如说3笔ABC，A申请了，那么返回的是BC的列表, 主定单不存在如果没有申请过退款，那么超时会挂在主定单上；比如ABC，返回主定单，ABC的超时和主定单相同     */
    TimeoutActionTime  *util.LocalTime `json:"timeout_action_time,omitempty" `

    /*
        交易快照地址     */
    SnapshotUrl  *string `json:"snapshot_url,omitempty" `

    /*
        分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)     */
    StepTradeStatus  *string `json:"step_trade_status,omitempty" `

    /*
        分阶段付款的已付金额（万人团订单已付金额）     */
    StepPaidFee  *string `json:"step_paid_fee,omitempty" `

    /*
        买家可以通过此字段查询是否当前交易可以评论，can_rate=true可以评价，false则不能评价。     */
    CanRate  *bool `json:"can_rate,omitempty" `

    /*
        卖家是否可以对订单进行评价     */
    SellerCanRate  *bool `json:"seller_can_rate,omitempty" `

    /*
        订单列表     */
    Orders  *[]TaobaoTradeSimpleGetOrder `json:"orders,omitempty" `

    /*
        优惠详情     */
    PromotionDetails  *[]TaobaoTradeSimpleGetPromotionDetails `json:"promotion_details,omitempty" `

    /*
        服务子订单列表     */
    ServiceOrders  *[]TaobaoTradeSimpleGetServiceOrders `json:"service_orders,omitempty" `

    /*
        交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔     */
    TradeFrom  *string `json:"trade_from,omitempty" `

    /*
        sellerNick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        购物金信息输出     */
    ExpandcardInfo  *TaobaoTradeSimpleGetExpandCardInfo `json:"expandcard_info,omitempty" `

    /*
        买卡订单本金     */
    ExpandCardBasicPrice  *string `json:"expand_card_basic_price,omitempty" `

    /*
        买卡订单权益金     */
    ExpandCardExpandPrice  *string `json:"expand_card_expand_price,omitempty" `

    /*
        用卡订单所用的本金     */
    ExpandCardBasicPriceUsed  *string `json:"expand_card_basic_price_used,omitempty" `

    /*
        用卡订单所用的权益金     */
    ExpandCardExpandPriceUsed  *string `json:"expand_card_expand_price_used,omitempty" `

    /*
        是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false     */
    IsPartConsign  *bool `json:"is_part_consign,omitempty" `

    /*
        处方药审核状态     */
    RxAuditStatus  *string `json:"rx_audit_status,omitempty" `

    /*
        使用红包付款金额     */
    CouponFee  *int64 `json:"coupon_fee,omitempty" `

    /*
        卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5     */
    SellerFlag  *int64 `json:"seller_flag,omitempty" `

    /*
        CRM系统中特有的ouid     */
    Ouid  *string `json:"ouid,omitempty" `

}

func (s *TaobaoTradeSimpleGetTrade) SetBuyerOpenUid(v string) *TaobaoTradeSimpleGetTrade {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetTitle(v string) *TaobaoTradeSimpleGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetType(v string) *TaobaoTradeSimpleGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetCreated(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetTid(v string) *TaobaoTradeSimpleGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSellerRate(v bool) *TaobaoTradeSimpleGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetBuyerRate(v bool) *TaobaoTradeSimpleGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetStatus(v string) *TaobaoTradeSimpleGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPayment(v string) *TaobaoTradeSimpleGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetDiscountFee(v string) *TaobaoTradeSimpleGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetAdjustFee(v string) *TaobaoTradeSimpleGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPostFee(v string) *TaobaoTradeSimpleGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetTotalFee(v string) *TaobaoTradeSimpleGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPayTime(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetEndTime(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetModified(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetConsignTime(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetReceivedPayment(v string) *TaobaoTradeSimpleGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetBuyerMemo(v string) *TaobaoTradeSimpleGetTrade {
    s.BuyerMemo = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSellerMemo(v string) *TaobaoTradeSimpleGetTrade {
    s.SellerMemo = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetBuyerMessage(v string) *TaobaoTradeSimpleGetTrade {
    s.BuyerMessage = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPicPath(v string) *TaobaoTradeSimpleGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetNumIid(v int64) *TaobaoTradeSimpleGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPrice(v string) *TaobaoTradeSimpleGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetShippingType(v string) *TaobaoTradeSimpleGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetNum(v int64) *TaobaoTradeSimpleGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetReceiverState(v string) *TaobaoTradeSimpleGetTrade {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetReceiverCity(v string) *TaobaoTradeSimpleGetTrade {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetAvailableConfirmFee(v string) *TaobaoTradeSimpleGetTrade {
    s.AvailableConfirmFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetHasPostFee(v bool) *TaobaoTradeSimpleGetTrade {
    s.HasPostFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetTimeoutActionTime(v util.LocalTime) *TaobaoTradeSimpleGetTrade {
    s.TimeoutActionTime = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSnapshotUrl(v string) *TaobaoTradeSimpleGetTrade {
    s.SnapshotUrl = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetStepTradeStatus(v string) *TaobaoTradeSimpleGetTrade {
    s.StepTradeStatus = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetStepPaidFee(v string) *TaobaoTradeSimpleGetTrade {
    s.StepPaidFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetCanRate(v bool) *TaobaoTradeSimpleGetTrade {
    s.CanRate = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSellerCanRate(v bool) *TaobaoTradeSimpleGetTrade {
    s.SellerCanRate = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetOrders(v []TaobaoTradeSimpleGetOrder) *TaobaoTradeSimpleGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetPromotionDetails(v []TaobaoTradeSimpleGetPromotionDetails) *TaobaoTradeSimpleGetTrade {
    s.PromotionDetails = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetServiceOrders(v []TaobaoTradeSimpleGetServiceOrders) *TaobaoTradeSimpleGetTrade {
    s.ServiceOrders = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetTradeFrom(v string) *TaobaoTradeSimpleGetTrade {
    s.TradeFrom = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSellerNick(v string) *TaobaoTradeSimpleGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetExpandcardInfo(v TaobaoTradeSimpleGetExpandCardInfo) *TaobaoTradeSimpleGetTrade {
    s.ExpandcardInfo = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetExpandCardBasicPrice(v string) *TaobaoTradeSimpleGetTrade {
    s.ExpandCardBasicPrice = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetExpandCardExpandPrice(v string) *TaobaoTradeSimpleGetTrade {
    s.ExpandCardExpandPrice = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetExpandCardBasicPriceUsed(v string) *TaobaoTradeSimpleGetTrade {
    s.ExpandCardBasicPriceUsed = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetExpandCardExpandPriceUsed(v string) *TaobaoTradeSimpleGetTrade {
    s.ExpandCardExpandPriceUsed = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetIsPartConsign(v bool) *TaobaoTradeSimpleGetTrade {
    s.IsPartConsign = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetRxAuditStatus(v string) *TaobaoTradeSimpleGetTrade {
    s.RxAuditStatus = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetCouponFee(v int64) *TaobaoTradeSimpleGetTrade {
    s.CouponFee = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetSellerFlag(v int64) *TaobaoTradeSimpleGetTrade {
    s.SellerFlag = &v
    return s
}
func (s *TaobaoTradeSimpleGetTrade) SetOuid(v string) *TaobaoTradeSimpleGetTrade {
    s.Ouid = &v
    return s
}
