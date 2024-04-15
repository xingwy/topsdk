package request

import (
        "topsdk/util"
    )

type TaobaoPromotionCouponsGetRequest struct {
    /*
        优惠券的id，唯一标识这个优惠券     */
    CouponId  *int64 `json:"coupon_id,omitempty" required:"false" `
    /*
        优惠券的截止日期     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" required:"false" `
    /*
        优惠券的面额，必须是3，5，10，20，50,100     */
    Denominations  *int64 `json:"denominations,omitempty" required:"false" `
    /*
        查询的页号，结果集是分页返回的，每页20条 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoPromotionCouponsGetRequest) SetCouponId(v int64) *TaobaoPromotionCouponsGetRequest {
    s.CouponId = &v
    return s
}
func (s *TaobaoPromotionCouponsGetRequest) SetEndTime(v util.LocalTime) *TaobaoPromotionCouponsGetRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoPromotionCouponsGetRequest) SetDenominations(v int64) *TaobaoPromotionCouponsGetRequest {
    s.Denominations = &v
    return s
}
func (s *TaobaoPromotionCouponsGetRequest) SetPageNo(v int64) *TaobaoPromotionCouponsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoPromotionCouponsGetRequest) SetPageSize(v int64) *TaobaoPromotionCouponsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoPromotionCouponsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CouponId != nil) {
        paramMap["coupon_id"] = *req.CouponId
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.Denominations != nil) {
        paramMap["denominations"] = *req.Denominations
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoPromotionCouponsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}