package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail struct {
	/*
	   限时打折名称     */
	LimitDiscountName *string `json:"limit_discount_name,omitempty" `

	/*
	   限时打折开始时间。     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   限时打折结束时间。     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   商品的id(数字类型)     */
	ItemId *int64 `json:"item_id,omitempty" `

	/*
	   该商品限时折扣     */
	ItemDiscount *string `json:"item_discount,omitempty" `

	/*
	   每人限购数量，1、2、5、10000(不限)。     */
	LimitNum *int64 `json:"limit_num,omitempty" `
}

func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetLimitDiscountName(v string) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.LimitDiscountName = &v
	return s
}
func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetStartTime(v util.LocalTime) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.StartTime = &v
	return s
}
func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetEndTime(v util.LocalTime) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.EndTime = &v
	return s
}
func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetItemId(v int64) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.ItemId = &v
	return s
}
func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetItemDiscount(v string) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.ItemDiscount = &v
	return s
}
func (s *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail) SetLimitNum(v int64) *TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail {
	s.LimitNum = &v
	return s
}
