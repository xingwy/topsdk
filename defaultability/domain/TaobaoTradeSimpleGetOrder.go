package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoTradeSimpleGetOrder struct {
	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图片的绝对路径     */
	PicPath *string `json:"pic_path,omitempty" `

	/*
	   商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	Price *string `json:"price,omitempty" `

	/*
	   商品数字ID     */
	NumIid *int64 `json:"num_iid,omitempty" `

	/*
	   商品的最小库存单位Sku的id.可以通过taobao.item.sku.get获取详细的Sku信息     */
	SkuId *string `json:"sku_id,omitempty" `

	/*
	   商家外部编码(可与商家外部系统对接)。外部商家自己定义的商品Item的id，可以通过taobao.items.custom.get获取商品的Item的信息     */
	OuterIid *string `json:"outer_iid,omitempty" `

	/*
	   外部网店自己定义的Sku编号     */
	OuterSkuId *string `json:"outer_sku_id,omitempty" `

	/*
	   退款状态。退款状态。可选值 WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
	RefundStatus *string `json:"refund_status,omitempty" `

	/*
	   订单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值: <ul><li>TRADE_NO_CREATE_PAY(没有创建支付宝交易) </li><li>WAIT_BUYER_PAY(等待买家付款) </li><li>WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) </li><li>WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) </li><li>TRADE_BUYER_SIGNED(买家已签收,货到付款专用) </li><li>TRADE_FINISHED(交易成功) </li><li>TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) </li><li>TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)</li><li>PAY_PENDING(国际信用卡支付付款确认中)</li></ul>     */
	Status *string `json:"status,omitempty" `

	/*
	   子订单编号     */
	Oid *int64 `json:"oid,omitempty" `

	/*
	   应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单级订单优惠金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	TotalFee *string `json:"total_fee,omitempty" `

	/*
	   子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。对于多子订单的交易，计算公式如下：payment = price * num + adjust_fee - discount_fee ；单子订单交易，payment与主订单的payment一致，对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。     */
	Payment *string `json:"payment,omitempty" `

	/*
	   子订单级订单优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
	DiscountFee *string `json:"discount_fee,omitempty" `

	/*
	   手工调整金额.格式为:1.01;单位:元;精确到小数点后两位.     */
	AdjustFee *string `json:"adjust_fee,omitempty" `

	/*
	   分摊之后的实付金额     */
	DivideOrderFee *string `json:"divide_order_fee,omitempty" `

	/*
	   优惠分摊     */
	PartMjzDiscount *string `json:"part_mjz_discount,omitempty" `

	/*
	   SKU的值。如：机身颜色:黑色;手机套餐:官方标配     */
	SkuPropertiesName *string `json:"sku_properties_name,omitempty" `

	/*
	   购买数量。取值范围:大于零的整数     */
	Num *int64 `json:"num,omitempty" `

	/*
	   订单快照URL     */
	SnapshotUrl *string `json:"snapshot_url,omitempty" `

	/*
	   订单超时到期时间。格式:yyyy-MM-dd HH:mm:ss     */
	TimeoutActionTime *util.LocalTime `json:"timeout_action_time,omitempty" `

	/*
	   买家是否已评价。可选值：true(已评价)，false(未评价)     */
	BuyerRate *bool `json:"buyer_rate,omitempty" `

	/*
	   卖家是否已评价。可选值：true(已评价)，false(未评价)     */
	SellerRate *bool `json:"seller_rate,omitempty" `

	/*
	   最近退款ID     */
	RefundId *string `json:"refund_id,omitempty" `

	/*
	   交易商品对应的类目ID     */
	Cid *int64 `json:"cid,omitempty" `

	/*
	   是否超卖     */
	IsOversold *bool `json:"is_oversold,omitempty" `

	/*
	   子订单的交易结束时间说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   子订单来源,如jhs(聚划算)、taobao(淘宝)、wap(无线)     */
	OrderFrom *string `json:"order_from,omitempty" `

	/*
	   子订单发货时间，当卖家对订单进行了多次发货，子订单的发货时间和主订单的发货时间可能不一样了，那么就需要以子订单的时间为准。（没有进行多次发货的订单，主订单的发货时间和子订单的发货时间都一样）     */
	ConsignTime *string `json:"consign_time,omitempty" `

	/*
	   子订单的运送方式（卖家对订单进行多次发货之后，一个主订单下的子订单的运送方式可能不同，用order.shipping_type来区分子订单的运送方式）     */
	ShippingType *string `json:"shipping_type,omitempty" `

	/*
	   子订单发货的快递公司名称     */
	LogisticsCompany *string `json:"logistics_company,omitempty" `

	/*
	   子订单所在包裹的运单号     */
	InvoiceNo *string `json:"invoice_no,omitempty" `

	/*
	   bind_oid字段的升级，支持返回绑定的多个子订单，多个子订单以半角逗号分隔     */
	BindOids *string `json:"bind_oids,omitempty" `

	/*
	   套餐ID     */
	ItemMealId *string `json:"item_meal_id,omitempty" `

	/*
	   购物金核销子订单权益金分摊金额（单位为元）     */
	ExpandCardExpandPriceUsedSuborder *string `json:"expand_card_expand_price_used_suborder,omitempty" `

	/*
	   购物金核销子订单本金分摊金额（单位为元）     */
	ExpandCardBasicPriceUsedSuborder *string `json:"expand_card_basic_price_used_suborder,omitempty" `

	/*
	   定制信息     */
	Customization *string `json:"customization,omitempty" `

	/*
	   是否为闲鱼订单 1是0否     */
	IsIdle *string `json:"is_idle,omitempty" `
}

func (s *TaobaoTradeSimpleGetOrder) SetTitle(v string) *TaobaoTradeSimpleGetOrder {
	s.Title = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetPicPath(v string) *TaobaoTradeSimpleGetOrder {
	s.PicPath = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetPrice(v string) *TaobaoTradeSimpleGetOrder {
	s.Price = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetNumIid(v int64) *TaobaoTradeSimpleGetOrder {
	s.NumIid = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetSkuId(v string) *TaobaoTradeSimpleGetOrder {
	s.SkuId = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetOuterIid(v string) *TaobaoTradeSimpleGetOrder {
	s.OuterIid = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetOuterSkuId(v string) *TaobaoTradeSimpleGetOrder {
	s.OuterSkuId = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetRefundStatus(v string) *TaobaoTradeSimpleGetOrder {
	s.RefundStatus = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetStatus(v string) *TaobaoTradeSimpleGetOrder {
	s.Status = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetOid(v int64) *TaobaoTradeSimpleGetOrder {
	s.Oid = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetTotalFee(v string) *TaobaoTradeSimpleGetOrder {
	s.TotalFee = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetPayment(v string) *TaobaoTradeSimpleGetOrder {
	s.Payment = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetDiscountFee(v string) *TaobaoTradeSimpleGetOrder {
	s.DiscountFee = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetAdjustFee(v string) *TaobaoTradeSimpleGetOrder {
	s.AdjustFee = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetDivideOrderFee(v string) *TaobaoTradeSimpleGetOrder {
	s.DivideOrderFee = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetPartMjzDiscount(v string) *TaobaoTradeSimpleGetOrder {
	s.PartMjzDiscount = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetSkuPropertiesName(v string) *TaobaoTradeSimpleGetOrder {
	s.SkuPropertiesName = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetNum(v int64) *TaobaoTradeSimpleGetOrder {
	s.Num = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetSnapshotUrl(v string) *TaobaoTradeSimpleGetOrder {
	s.SnapshotUrl = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetTimeoutActionTime(v util.LocalTime) *TaobaoTradeSimpleGetOrder {
	s.TimeoutActionTime = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetBuyerRate(v bool) *TaobaoTradeSimpleGetOrder {
	s.BuyerRate = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetSellerRate(v bool) *TaobaoTradeSimpleGetOrder {
	s.SellerRate = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetRefundId(v string) *TaobaoTradeSimpleGetOrder {
	s.RefundId = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetCid(v int64) *TaobaoTradeSimpleGetOrder {
	s.Cid = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetIsOversold(v bool) *TaobaoTradeSimpleGetOrder {
	s.IsOversold = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetEndTime(v util.LocalTime) *TaobaoTradeSimpleGetOrder {
	s.EndTime = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetOrderFrom(v string) *TaobaoTradeSimpleGetOrder {
	s.OrderFrom = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetConsignTime(v string) *TaobaoTradeSimpleGetOrder {
	s.ConsignTime = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetShippingType(v string) *TaobaoTradeSimpleGetOrder {
	s.ShippingType = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetLogisticsCompany(v string) *TaobaoTradeSimpleGetOrder {
	s.LogisticsCompany = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetInvoiceNo(v string) *TaobaoTradeSimpleGetOrder {
	s.InvoiceNo = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetBindOids(v string) *TaobaoTradeSimpleGetOrder {
	s.BindOids = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetItemMealId(v string) *TaobaoTradeSimpleGetOrder {
	s.ItemMealId = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetExpandCardExpandPriceUsedSuborder(v string) *TaobaoTradeSimpleGetOrder {
	s.ExpandCardExpandPriceUsedSuborder = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetExpandCardBasicPriceUsedSuborder(v string) *TaobaoTradeSimpleGetOrder {
	s.ExpandCardBasicPriceUsedSuborder = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetCustomization(v string) *TaobaoTradeSimpleGetOrder {
	s.Customization = &v
	return s
}
func (s *TaobaoTradeSimpleGetOrder) SetIsIdle(v string) *TaobaoTradeSimpleGetOrder {
	s.IsIdle = &v
	return s
}
