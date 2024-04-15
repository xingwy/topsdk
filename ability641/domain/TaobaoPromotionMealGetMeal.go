package domain


type TaobaoPromotionMealGetMeal struct {
    /*
        套餐id。     */
    MealId  *int64 `json:"meal_id,omitempty" `

    /*
        搭配套餐名称。     */
    MealName  *string `json:"meal_name,omitempty" `

    /*
        套餐一口价(单位是：分)     */
    MealPrice  *string `json:"meal_price,omitempty" `

    /*
        搭配套餐商品列表。item_id为商品的id;item_show_name为商品显示名。因最多允许5个商品进行搭配，所以查询最多有5个，以json格式传出。     */
    ItemList  *string `json:"item_list,omitempty" `

    /*
        运费模板类型。卖家标识'SELL';买家标识'BUY'。若为'SELL'则字段postage_id为空。若为'BUY'，则postage_id为运费模板id，必有值。     */
    TypePostage  *string `json:"type_postage,omitempty" `

    /*
        搭配套餐描述！     */
    MealMemo  *string `json:"meal_memo,omitempty" `

    /*
        普通运费模板id。若这个字段为空或0时，运费是卖家负责;若这个字段不为空，说明运费模板存在，运费是买家负责。     */
    PostageId  *int64 `json:"postage_id,omitempty" `

    /*
        套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoPromotionMealGetMeal) SetMealId(v int64) *TaobaoPromotionMealGetMeal {
    s.MealId = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetMealName(v string) *TaobaoPromotionMealGetMeal {
    s.MealName = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetMealPrice(v string) *TaobaoPromotionMealGetMeal {
    s.MealPrice = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetItemList(v string) *TaobaoPromotionMealGetMeal {
    s.ItemList = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetTypePostage(v string) *TaobaoPromotionMealGetMeal {
    s.TypePostage = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetMealMemo(v string) *TaobaoPromotionMealGetMeal {
    s.MealMemo = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetPostageId(v int64) *TaobaoPromotionMealGetMeal {
    s.PostageId = &v
    return s
}
func (s *TaobaoPromotionMealGetMeal) SetStatus(v string) *TaobaoPromotionMealGetMeal {
    s.Status = &v
    return s
}
