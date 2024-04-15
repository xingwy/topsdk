package domain


type TaobaoItemsSellerListGetSkuFeature struct {
    /*
        colorHotNew     */
    ColorHotNew  *string `json:"color_hot_new,omitempty" `

    /*
        colorMaterialImg     */
    ColorMaterialImg  *string `json:"color_material_img,omitempty" `

    /*
        colorValue     */
    ColorValue  *string `json:"color_value,omitempty" `

    /*
        colorMaterial     */
    ColorMaterial  *string `json:"color_material,omitempty" `

    /*
        colorSeries     */
    ColorSeries  *string `json:"color_series,omitempty" `

}

func (s *TaobaoItemsSellerListGetSkuFeature) SetColorHotNew(v string) *TaobaoItemsSellerListGetSkuFeature {
    s.ColorHotNew = &v
    return s
}
func (s *TaobaoItemsSellerListGetSkuFeature) SetColorMaterialImg(v string) *TaobaoItemsSellerListGetSkuFeature {
    s.ColorMaterialImg = &v
    return s
}
func (s *TaobaoItemsSellerListGetSkuFeature) SetColorValue(v string) *TaobaoItemsSellerListGetSkuFeature {
    s.ColorValue = &v
    return s
}
func (s *TaobaoItemsSellerListGetSkuFeature) SetColorMaterial(v string) *TaobaoItemsSellerListGetSkuFeature {
    s.ColorMaterial = &v
    return s
}
func (s *TaobaoItemsSellerListGetSkuFeature) SetColorSeries(v string) *TaobaoItemsSellerListGetSkuFeature {
    s.ColorSeries = &v
    return s
}
