package response

import (
	"github.com/xingwy/topsdk/ability641/domain"
)

type TaobaoPromotionMealGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   搭配套餐列表。
	*/
	MealList []domain.TaobaoPromotionMealGetMeal `json:"meal_list,omitempty" `
}
