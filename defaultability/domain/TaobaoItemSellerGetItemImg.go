package domain


type TaobaoItemSellerGetItemImg struct {
    /*
        商品图片的id，和商品相对应（主图id默认为0）     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片放在第几张（多图时可设置）     */
    Position  *int64 `json:"position,omitempty" `

    /*
        图片链接地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoItemSellerGetItemImg) SetId(v int64) *TaobaoItemSellerGetItemImg {
    s.Id = &v
    return s
}
func (s *TaobaoItemSellerGetItemImg) SetPosition(v int64) *TaobaoItemSellerGetItemImg {
    s.Position = &v
    return s
}
func (s *TaobaoItemSellerGetItemImg) SetUrl(v string) *TaobaoItemSellerGetItemImg {
    s.Url = &v
    return s
}
