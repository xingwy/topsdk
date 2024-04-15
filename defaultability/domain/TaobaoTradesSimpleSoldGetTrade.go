package domain


type TaobaoTradesSimpleSoldGetTrade struct {
    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

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
    Created  *string `json:"created,omitempty" `

    /*
        交易编号 (父订单的交易编号)     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        物流运单号     */
    Sid  *string `json:"sid,omitempty" `

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
    PayTime  *string `json:"pay_time,omitempty" `

    /*
        交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss     */
    EndTime  *string `json:"end_time,omitempty" `

    /*
        交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        同步到卖家库的时间，taobao.trades.sold.incrementv.get接口返回此字段     */
    AsyncModified  *string `json:"async_modified,omitempty" `

    /*
        卖家发货时间。格式:yyyy-MM-dd HH:mm:ss     */
    ConsignTime  *string `json:"consign_time,omitempty" `

    /*
        卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    ReceivedPayment  *string `json:"received_payment,omitempty" `

    /*
        交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    CommissionFee  *string `json:"commission_fee,omitempty" `

    /*
        买家下单的地区     */
    BuyerArea  *string `json:"buyer_area,omitempty" `

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
        买家使用积分,下单时生成，且一直不变。格式:100;单位:个.     */
    PointFee  *int64 `json:"point_fee,omitempty" `

    /*
        买家实际使用积分（扣除部分退款使用的积分），交易完成后生成（交易成功或关闭），交易未完成时该字段值为0。格式:100;单位:个     */
    RealPointFee  *int64 `json:"real_point_fee,omitempty" `

    /*
        买家获得积分,返点的积分。格式:100;单位:个。返点的积分要交易成功之后才能获得。     */
    BuyerObtainPointFee  *int64 `json:"buyer_obtain_point_fee,omitempty" `

    /*
        收货人的所在省份     */
    ReceiverState  *string `json:"receiver_state,omitempty" `

    /*
        收货人的所在城市<br>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5     */
    SellerFlag  *int64 `json:"seller_flag,omitempty" `

    /*
        判断订单是否有买家留言，有买家留言返回true，否则返回false     */
    HasBuyerMessage  *bool `json:"has_buyer_message,omitempty" `

    /*
        交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔     */
    TradeFrom  *string `json:"trade_from,omitempty" `

    /*
        分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)     */
    StepTradeStatus  *string `json:"step_trade_status,omitempty" `

    /*
        分阶段付款的已付金额（万人团订单已付金额）     */
    StepPaidFee  *string `json:"step_paid_fee,omitempty" `

    /*
        交易外部来源：ownshop(商家官网)     */
    TradeSource  *string `json:"trade_source,omitempty" `

    /*
        订单将在此时间前发出，主要用于预售订单     */
    SendTime  *string `json:"send_time,omitempty" `

    /*
        卖家是否可以对订单进行评价     */
    SellerCanRate  *bool `json:"seller_can_rate,omitempty" `

    /*
        订单列表     */
    Orders  *[]TaobaoTradesSimpleSoldGetOrder `json:"orders,omitempty" `

    /*
        服务子订单列表     */
    ServiceOrders  *[]TaobaoTradesSimpleSoldGetServiceOrder `json:"service_orders,omitempty" `

    /*
        处方药未审核     */
    RxAuditStatus  *string `json:"rx_audit_status,omitempty" `

    /*
        是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false     */
    IsPartConsign  *bool `json:"is_part_consign,omitempty" `

    /*
        CRM系统中特有的ouid     */
    Ouid  *string `json:"ouid,omitempty" `

}

func (s *TaobaoTradesSimpleSoldGetTrade) SetSellerNick(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetBuyerOpenUid(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetTitle(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetType(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetCreated(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetTid(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetSid(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Sid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetSellerRate(v bool) *TaobaoTradesSimpleSoldGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetBuyerRate(v bool) *TaobaoTradesSimpleSoldGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetStatus(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPayment(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetDiscountFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetAdjustFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPostFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetTotalFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPayTime(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetEndTime(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetModified(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetAsyncModified(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.AsyncModified = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetConsignTime(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetReceivedPayment(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetCommissionFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetBuyerArea(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.BuyerArea = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPicPath(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetNumIid(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPrice(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetShippingType(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetNum(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetPointFee(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.PointFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetRealPointFee(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.RealPointFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetBuyerObtainPointFee(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetReceiverState(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetReceiverCity(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetSellerFlag(v int64) *TaobaoTradesSimpleSoldGetTrade {
    s.SellerFlag = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetHasBuyerMessage(v bool) *TaobaoTradesSimpleSoldGetTrade {
    s.HasBuyerMessage = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetTradeFrom(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.TradeFrom = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetStepTradeStatus(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.StepTradeStatus = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetStepPaidFee(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.StepPaidFee = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetTradeSource(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.TradeSource = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetSendTime(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.SendTime = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetSellerCanRate(v bool) *TaobaoTradesSimpleSoldGetTrade {
    s.SellerCanRate = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetOrders(v []TaobaoTradesSimpleSoldGetOrder) *TaobaoTradesSimpleSoldGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetServiceOrders(v []TaobaoTradesSimpleSoldGetServiceOrder) *TaobaoTradesSimpleSoldGetTrade {
    s.ServiceOrders = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetRxAuditStatus(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.RxAuditStatus = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetIsPartConsign(v bool) *TaobaoTradesSimpleSoldGetTrade {
    s.IsPartConsign = &v
    return s
}
func (s *TaobaoTradesSimpleSoldGetTrade) SetOuid(v string) *TaobaoTradesSimpleSoldGetTrade {
    s.Ouid = &v
    return s
}
