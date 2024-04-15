package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoVasOrderSearchArticleBizOrder struct {
	/*
	   淘宝会员名     */
	Nick *string `json:"nick,omitempty" `

	/*
	   应用名称     */
	ArticleName *string `json:"article_name,omitempty" `

	/*
	   应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码     */
	ArticleCode *string `json:"article_code,omitempty" `

	/*
	   商品模型名称     */
	ArticleItemName *string `json:"article_item_name,omitempty" `

	/*
	   收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码     */
	ItemCode *string `json:"item_code,omitempty" `

	/*
	   订单创建时间（订购时间）     */
	Create *util.LocalTime `json:"create,omitempty" `

	/*
	   订购周期  1表示年 ，2表示月，3表示天。     */
	OrderCycle *string `json:"order_cycle,omitempty" `

	/*
	   订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到）     */
	BizType *int64 `json:"biz_type,omitempty" `

	/*
	   子订单号     */
	OrderId *int64 `json:"order_id,omitempty" `

	/*
	   原价（单位为分）     */
	Fee *string `json:"fee,omitempty" `

	/*
	   优惠（单位为分）     */
	PromFee *string `json:"prom_fee,omitempty" `

	/*
	   实付（单位为分）     */
	TotalPayFee *string `json:"total_pay_fee,omitempty" `

	/*
	   订购周期开始时间     */
	OrderCycleStart *util.LocalTime `json:"order_cycle_start,omitempty" `

	/*
	   订购周期结束时间     */
	OrderCycleEnd *util.LocalTime `json:"order_cycle_end,omitempty" `

	/*
	   退款（单位为分；升级时，系统会将升级前老版本按照剩余订购天数退还剩余金额）     */
	RefundFee *string `json:"refund_fee,omitempty" `

	/*
	   订单号     */
	BizOrderId *int64 `json:"biz_order_id,omitempty" `

	/*
	   activityCode     */
	ActivityCode *string `json:"activity_code,omitempty" `

	/*
	   业务订单状态     */
	OrderBizStatus *string `json:"order_biz_status,omitempty" `
}

func (s *TaobaoVasOrderSearchArticleBizOrder) SetNick(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.Nick = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetArticleName(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.ArticleName = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetArticleCode(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.ArticleCode = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetArticleItemName(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.ArticleItemName = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetItemCode(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.ItemCode = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetCreate(v util.LocalTime) *TaobaoVasOrderSearchArticleBizOrder {
	s.Create = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetOrderCycle(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.OrderCycle = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetBizType(v int64) *TaobaoVasOrderSearchArticleBizOrder {
	s.BizType = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetOrderId(v int64) *TaobaoVasOrderSearchArticleBizOrder {
	s.OrderId = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetFee(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.Fee = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetPromFee(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.PromFee = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetTotalPayFee(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.TotalPayFee = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetOrderCycleStart(v util.LocalTime) *TaobaoVasOrderSearchArticleBizOrder {
	s.OrderCycleStart = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetOrderCycleEnd(v util.LocalTime) *TaobaoVasOrderSearchArticleBizOrder {
	s.OrderCycleEnd = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetRefundFee(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.RefundFee = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetBizOrderId(v int64) *TaobaoVasOrderSearchArticleBizOrder {
	s.BizOrderId = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetActivityCode(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.ActivityCode = &v
	return s
}
func (s *TaobaoVasOrderSearchArticleBizOrder) SetOrderBizStatus(v string) *TaobaoVasOrderSearchArticleBizOrder {
	s.OrderBizStatus = &v
	return s
}
