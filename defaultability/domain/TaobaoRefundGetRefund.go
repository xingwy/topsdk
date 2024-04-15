package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundGetRefund struct {
	/*
	   卖家收货地址     */
	Address *string `json:"address,omitempty" `

	/*
	   退款先行垫付默认的未申请状态 0;退款先行垫付申请中  1;退款先行垫付，垫付完成 2;退款先行垫付，卖家拒绝收货 3;退款先行垫付，垫付关闭 4;退款先行垫付，垫付分账成功 5;     */
	AdvanceStatus *int64 `json:"advance_status,omitempty" `

	/*
	   支付宝交易号     */
	AlipayNo *string `json:"alipay_no,omitempty" `

	/*
	   attribute     */
	Attribute *string `json:"attribute,omitempty" `

	/*
	   买家昵称     */
	BuyerNick *string `json:"buyer_nick,omitempty" `

	/*
	   物流公司名称     */
	CompanyName *string `json:"company_name,omitempty" `

	/*
	   退款申请时间。格式:yyyy-MM-dd HH:mm:ss     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   不需客服介入1;需要客服介入2;客服已经介入3;客服初审完成4;客服主管复审失败5;客服处理完成6;系统撤销(B2B使用)，维权撤销(集市使用) 7;支持买家 8;支持卖家 9;举证中 10;开放申诉 11;     */
	CsStatus *int64 `json:"cs_status,omitempty" `

	/*
	   退款说明     */
	Desc *string `json:"desc,omitempty" `

	/*
	   退货时间。格式:yyyy-MM-dd HH:mm:ss     */
	GoodReturnTime *util.LocalTime `json:"good_return_time,omitempty" `

	/*
	   货物状态。可选值BUYER_NOT_RECEIVED (买家未收到货) BUYER_RECEIVED (买家已收到货) BUYER_RETURNED_GOODS (买家已退货)     */
	GoodStatus *string `json:"good_status,omitempty" `

	/*
	   买家是否需要退货。可选值:true(是),false(否)     */
	HasGoodReturn *bool `json:"has_good_return,omitempty" `

	/*
	   更新时间。格式:yyyy-MM-dd HH:mm:ss     */
	Modified *util.LocalTime `json:"modified,omitempty" `

	/*
	   商品数量     */
	Num *int64 `json:"num,omitempty" `

	/*
	   申请退款的商品数字编号     */
	NumIid *int64 `json:"num_iid,omitempty" `

	/*
	   子订单号。如果是单笔交易oid会等于tid     */
	Oid *int64 `json:"oid,omitempty" `

	/*
	   退款约束，可选值：cannot_refuse（不允许操作），refund_onweb（需要到网页版操作）     */
	OperationContraint *string `json:"operation_contraint,omitempty" `

	/*
	   退款对应的订单交易状态。可选值TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) 取自"http://open.taobao.com/dev/index.php/%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81"     */
	OrderStatus *string `json:"order_status,omitempty" `

	/*
	   商品外部商家编码     */
	OuterId *string `json:"outer_id,omitempty" `

	/*
	   支付给卖家的金额(交易总金额-退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	Payment *string `json:"payment,omitempty" `

	/*
	   商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	Price *string `json:"price,omitempty" `

	/*
	   退款原因     */
	Reason *string `json:"reason,omitempty" `

	/*
	   退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	RefundFee *string `json:"refund_fee,omitempty" `

	/*
	   退款单号     */
	RefundId *string `json:"refund_id,omitempty" `

	/*
	   退款阶段，可选值：onsale(售中)/aftersale(售后)     */
	RefundPhase *string `json:"refund_phase,omitempty" `

	/*
	   退款超时结构RefundRemindTimeout     */
	RefundRemindTimeout *TaobaoRefundGetRefundRemindTimeout `json:"refund_remind_timeout,omitempty" `

	/*
	   退款版本号（时间戳）     */
	RefundVersion *int64 `json:"refund_version,omitempty" `

	/*
	   卖家昵称     */
	SellerNick *string `json:"seller_nick,omitempty" `

	/*
	   物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS).     */
	ShippingType *string `json:"shipping_type,omitempty" `

	/*
	   退货运单号     */
	Sid *string `json:"sid,omitempty" `

	/*
	   商品SKU信息     */
	Sku *string `json:"sku,omitempty" `

	/*
	   分账给卖家的钱     */
	SplitSellerFee *string `json:"split_seller_fee,omitempty" `

	/*
	   分账给淘宝的钱     */
	SplitTaobaoFee *string `json:"split_taobao_fee,omitempty" `

	/*
	   退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
	Status *string `json:"status,omitempty" `

	/*
	   淘宝交易单号     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   交易总金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	TotalFee *string `json:"total_fee,omitempty" `

	/*
	   买家账号的OpenUID     */
	BuyerOpenUid *string `json:"buyer_open_uid,omitempty" `

	/*
	   crm改造新增ouid返回     */
	Ouid *string `json:"ouid,omitempty" `

	/*
	   组合品信息     */
	CombineItemInfo *[]TaobaoRefundGetCombineSubItemDO `json:"combine_item_info,omitempty" `

	/*
	   退款类型，可选值REFUND(仅退款),REFUND_AND_RETURN(退货退款),TMALL_EXCHANGE(天猫换货),TAOBAO_EXCHANGE(淘宝换货),REPAIR(维修),RESHIPPING(补寄),OTHERS(其他)     */
	DisputeType *string `json:"dispute_type,omitempty" `

	/*
	   完结时间。格式:yyyy-MM-dd HH:mm:ss     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   卖家账号的OpenUID     */
	SellerOpenUid *string `json:"seller_open_uid,omitempty" `
}

func (s *TaobaoRefundGetRefund) SetAddress(v string) *TaobaoRefundGetRefund {
	s.Address = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetAdvanceStatus(v int64) *TaobaoRefundGetRefund {
	s.AdvanceStatus = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetAlipayNo(v string) *TaobaoRefundGetRefund {
	s.AlipayNo = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetAttribute(v string) *TaobaoRefundGetRefund {
	s.Attribute = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetBuyerNick(v string) *TaobaoRefundGetRefund {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetCompanyName(v string) *TaobaoRefundGetRefund {
	s.CompanyName = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetCreated(v util.LocalTime) *TaobaoRefundGetRefund {
	s.Created = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetCsStatus(v int64) *TaobaoRefundGetRefund {
	s.CsStatus = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetDesc(v string) *TaobaoRefundGetRefund {
	s.Desc = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetGoodReturnTime(v util.LocalTime) *TaobaoRefundGetRefund {
	s.GoodReturnTime = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetGoodStatus(v string) *TaobaoRefundGetRefund {
	s.GoodStatus = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetHasGoodReturn(v bool) *TaobaoRefundGetRefund {
	s.HasGoodReturn = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetModified(v util.LocalTime) *TaobaoRefundGetRefund {
	s.Modified = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetNum(v int64) *TaobaoRefundGetRefund {
	s.Num = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetNumIid(v int64) *TaobaoRefundGetRefund {
	s.NumIid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetOid(v int64) *TaobaoRefundGetRefund {
	s.Oid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetOperationContraint(v string) *TaobaoRefundGetRefund {
	s.OperationContraint = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetOrderStatus(v string) *TaobaoRefundGetRefund {
	s.OrderStatus = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetOuterId(v string) *TaobaoRefundGetRefund {
	s.OuterId = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetPayment(v string) *TaobaoRefundGetRefund {
	s.Payment = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetPrice(v string) *TaobaoRefundGetRefund {
	s.Price = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetReason(v string) *TaobaoRefundGetRefund {
	s.Reason = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetRefundFee(v string) *TaobaoRefundGetRefund {
	s.RefundFee = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetRefundId(v string) *TaobaoRefundGetRefund {
	s.RefundId = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetRefundPhase(v string) *TaobaoRefundGetRefund {
	s.RefundPhase = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetRefundRemindTimeout(v TaobaoRefundGetRefundRemindTimeout) *TaobaoRefundGetRefund {
	s.RefundRemindTimeout = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetRefundVersion(v int64) *TaobaoRefundGetRefund {
	s.RefundVersion = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSellerNick(v string) *TaobaoRefundGetRefund {
	s.SellerNick = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetShippingType(v string) *TaobaoRefundGetRefund {
	s.ShippingType = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSid(v string) *TaobaoRefundGetRefund {
	s.Sid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSku(v string) *TaobaoRefundGetRefund {
	s.Sku = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSplitSellerFee(v string) *TaobaoRefundGetRefund {
	s.SplitSellerFee = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSplitTaobaoFee(v string) *TaobaoRefundGetRefund {
	s.SplitTaobaoFee = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetStatus(v string) *TaobaoRefundGetRefund {
	s.Status = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetTid(v int64) *TaobaoRefundGetRefund {
	s.Tid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetTitle(v string) *TaobaoRefundGetRefund {
	s.Title = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetTotalFee(v string) *TaobaoRefundGetRefund {
	s.TotalFee = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetBuyerOpenUid(v string) *TaobaoRefundGetRefund {
	s.BuyerOpenUid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetOuid(v string) *TaobaoRefundGetRefund {
	s.Ouid = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetCombineItemInfo(v []TaobaoRefundGetCombineSubItemDO) *TaobaoRefundGetRefund {
	s.CombineItemInfo = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetDisputeType(v string) *TaobaoRefundGetRefund {
	s.DisputeType = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetEndTime(v util.LocalTime) *TaobaoRefundGetRefund {
	s.EndTime = &v
	return s
}
func (s *TaobaoRefundGetRefund) SetSellerOpenUid(v string) *TaobaoRefundGetRefund {
	s.SellerOpenUid = &v
	return s
}
