package domain


type TaobaoItemcatsAuthorizeGetSellerAuthorize struct {
    /*
        商品类目结构     */
    ItemCats  *[]TaobaoItemcatsAuthorizeGetItemCat `json:"item_cats,omitempty" `

    /*
        商品类目结构     */
    XinpinItemCats  *[]TaobaoItemcatsAuthorizeGetItemCat `json:"xinpin_item_cats,omitempty" `

    /*
        品牌     */
    Brands  *[]TaobaoItemcatsAuthorizeGetBrand `json:"brands,omitempty" `

}

func (s *TaobaoItemcatsAuthorizeGetSellerAuthorize) SetItemCats(v []TaobaoItemcatsAuthorizeGetItemCat) *TaobaoItemcatsAuthorizeGetSellerAuthorize {
    s.ItemCats = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetSellerAuthorize) SetXinpinItemCats(v []TaobaoItemcatsAuthorizeGetItemCat) *TaobaoItemcatsAuthorizeGetSellerAuthorize {
    s.XinpinItemCats = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetSellerAuthorize) SetBrands(v []TaobaoItemcatsAuthorizeGetBrand) *TaobaoItemcatsAuthorizeGetSellerAuthorize {
    s.Brands = &v
    return s
}
