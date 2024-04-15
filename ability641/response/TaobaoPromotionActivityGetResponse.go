package response

import (
	"github.com/xingwy/topsdk/ability641/domain"
)

type TaobaoPromotionActivityGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   活动列表
	*/
	Activitys []domain.TaobaoPromotionActivityGetActivity `json:"activitys,omitempty" `
}
