package request


type TaobaoPromotionMealGetRequest struct {
    /*
        搭配套餐id     */
    MealId  *int64 `json:"meal_id,omitempty" required:"false" `
    /*
        套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。     */
    Status  *string `json:"status,omitempty" required:"false" `
}

func (s *TaobaoPromotionMealGetRequest) SetMealId(v int64) *TaobaoPromotionMealGetRequest {
    s.MealId = &v
    return s
}
func (s *TaobaoPromotionMealGetRequest) SetStatus(v string) *TaobaoPromotionMealGetRequest {
    s.Status = &v
    return s
}

func (req *TaobaoPromotionMealGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MealId != nil) {
        paramMap["meal_id"] = *req.MealId
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoPromotionMealGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}