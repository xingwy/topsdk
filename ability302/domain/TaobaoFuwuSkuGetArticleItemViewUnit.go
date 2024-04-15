package domain


type TaobaoFuwuSkuGetArticleItemViewUnit struct {
    /*
        需要支付的价格，单位：元     */
    ActualPrice  *string `json:"actual_price,omitempty" `

    /*
        用户是否可以购买     */
    CanSub  *bool `json:"can_sub,omitempty" `

    /*
        周期数，如1，3，6，12。对于周期型和周期计量型返回。     */
    CycNum  *int64 `json:"cyc_num,omitempty" `

    /*
        1-年，2-月，3-日。对于周期型和周期计量型返回。     */
    CycUnit  *int64 `json:"cyc_unit,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误文案     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        收费项目code     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        收费项目名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        原价，单位：元     */
    OriginPrice  *string `json:"origin_price,omitempty" `

    /*
        优惠，单位：元     */
    PromPrice  *string `json:"prom_price,omitempty" `

    /*
        数量。对于周期计量型返回计量数。     */
    Quantity  *int64 `json:"quantity,omitempty" `

}

func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetActualPrice(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.ActualPrice = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetCanSub(v bool) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.CanSub = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetCycNum(v int64) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.CycNum = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetCycUnit(v int64) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.CycUnit = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetErrorCode(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetErrorMsg(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetItemCode(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.ItemCode = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetItemName(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.ItemName = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetOriginPrice(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.OriginPrice = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetPromPrice(v string) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.PromPrice = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleItemViewUnit) SetQuantity(v int64) *TaobaoFuwuSkuGetArticleItemViewUnit {
    s.Quantity = &v
    return s
}
