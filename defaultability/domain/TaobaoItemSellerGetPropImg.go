package domain


type TaobaoItemSellerGetPropImg struct {
    /*
        属性图片的id，和商品相对应     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片放在第几张（多图时可设置）     */
    Position  *int64 `json:"position,omitempty" `

    /*
        图片所对应的属性组合的字符串     */
    Properties  *string `json:"properties,omitempty" `

    /*
        图片链接地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoItemSellerGetPropImg) SetId(v int64) *TaobaoItemSellerGetPropImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemSellerGetPropImg) SetPosition(v int64) *TaobaoItemSellerGetPropImg {
    s.Position = &v
    return s
}
func (s *TaobaoItemSellerGetPropImg) SetProperties(v string) *TaobaoItemSellerGetPropImg {
    s.Properties = &v
    return s
}
func (s *TaobaoItemSellerGetPropImg) SetUrl(v string) *TaobaoItemSellerGetPropImg {
    s.Url = &v
    return s
}
