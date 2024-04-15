package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoCrmGradeGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   等级信息集合
	*/
	GradePromotions []domain.TaobaoCrmGradeGetGradePromotion `json:"grade_promotions,omitempty" `
}
