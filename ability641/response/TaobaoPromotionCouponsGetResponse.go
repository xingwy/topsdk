package response

import (
    "topsdk/ability641/domain"
)

type TaobaoPromotionCouponsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询的总数量
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        优惠券列表
    */
    Coupons  []domain.TaobaoPromotionCouponsGetCoupon `json:"coupons,omitempty" `
}
