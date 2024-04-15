package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoPromotionLimitdiscountDetailGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   限时打折对应的商品详情列表。
	*/
	ItemDiscountDetailList []domain.TaobaoPromotionLimitdiscountDetailGetLimitDiscountDetail `json:"item_discount_detail_list,omitempty" `
}
