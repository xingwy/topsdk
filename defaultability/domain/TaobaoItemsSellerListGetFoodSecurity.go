package domain


type TaobaoItemsSellerListGetFoodSecurity struct {
    /*
        厂家联系方式     */
    Contact  *string `json:"contact,omitempty" `

    /*
        产品标准号     */
    DesignCode  *string `json:"design_code,omitempty" `

    /*
        厂名     */
    Factory  *string `json:"factory,omitempty" `

    /*
        厂址     */
    FactorySite  *string `json:"factory_site,omitempty" `

    /*
        食品添加剂     */
    FoodAdditive  *string `json:"food_additive,omitempty" `

    /*
        健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；     */
    HealthProductNo  *string `json:"health_product_no,omitempty" `

    /*
        配料表     */
    Mix  *string `json:"mix,omitempty" `

    /*
        保质期     */
    Period  *string `json:"period,omitempty" `

    /*
        储藏方法     */
    PlanStorage  *string `json:"plan_storage,omitempty" `

    /*
        生产许可证号     */
    PrdLicenseNo  *string `json:"prd_license_no,omitempty" `

    /*
        生产结束日期     */
    ProductDateEnd  *string `json:"product_date_end,omitempty" `

    /*
        生产开始日期     */
    ProductDateStart  *string `json:"product_date_start,omitempty" `

    /*
        进货结束日期，要在生产日期之后     */
    StockDateEnd  *string `json:"stock_date_end,omitempty" `

    /*
        进货开始日期，要在生产日期之后     */
    StockDateStart  *string `json:"stock_date_start,omitempty" `

    /*
        供货商     */
    Supplier  *string `json:"supplier,omitempty" `

}

func (s *TaobaoItemsSellerListGetFoodSecurity) SetContact(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.Contact = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetDesignCode(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.DesignCode = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetFactory(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.Factory = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetFactorySite(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.FactorySite = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetFoodAdditive(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.FoodAdditive = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetHealthProductNo(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.HealthProductNo = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetMix(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.Mix = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetPeriod(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.Period = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetPlanStorage(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.PlanStorage = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetPrdLicenseNo(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.PrdLicenseNo = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetProductDateEnd(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.ProductDateEnd = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetProductDateStart(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.ProductDateStart = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetStockDateEnd(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.StockDateEnd = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetStockDateStart(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.StockDateStart = &v
    return s
}
func (s *TaobaoItemsSellerListGetFoodSecurity) SetSupplier(v string) *TaobaoItemsSellerListGetFoodSecurity {
    s.Supplier = &v
    return s
}
