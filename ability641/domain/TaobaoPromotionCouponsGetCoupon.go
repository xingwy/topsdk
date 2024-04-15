package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoPromotionCouponsGetCoupon struct {
	/*
	   优惠券ID     */
	CouponId *int64 `json:"coupon_id,omitempty" `

	/*
	   优惠券的面值，返回的是“分”，不是“元”，500代表500分相当于5元     */
	Denominations *int64 `json:"denominations,omitempty" `

	/*
	   优惠券创建时间     */
	CreatTime *util.LocalTime `json:"creat_time,omitempty" `

	/*
	   优惠券的截止日期     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   订单满多少分才能用这个优惠券，501就是满501分能使用。注意：返回的是“分”，不是“元”     */
	Condition *int64 `json:"condition,omitempty" `

	/*
	   优惠券的创建渠道，自己创建/他人创建     */
	CreateChannel *string `json:"create_channel,omitempty" `
}

func (s *TaobaoPromotionCouponsGetCoupon) SetCouponId(v int64) *TaobaoPromotionCouponsGetCoupon {
	s.CouponId = &v
	return s
}
func (s *TaobaoPromotionCouponsGetCoupon) SetDenominations(v int64) *TaobaoPromotionCouponsGetCoupon {
	s.Denominations = &v
	return s
}
func (s *TaobaoPromotionCouponsGetCoupon) SetCreatTime(v util.LocalTime) *TaobaoPromotionCouponsGetCoupon {
	s.CreatTime = &v
	return s
}
func (s *TaobaoPromotionCouponsGetCoupon) SetEndTime(v util.LocalTime) *TaobaoPromotionCouponsGetCoupon {
	s.EndTime = &v
	return s
}
func (s *TaobaoPromotionCouponsGetCoupon) SetCondition(v int64) *TaobaoPromotionCouponsGetCoupon {
	s.Condition = &v
	return s
}
func (s *TaobaoPromotionCouponsGetCoupon) SetCreateChannel(v string) *TaobaoPromotionCouponsGetCoupon {
	s.CreateChannel = &v
	return s
}
