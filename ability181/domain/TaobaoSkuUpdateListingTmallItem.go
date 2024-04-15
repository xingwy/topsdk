package domain


type TaobaoSkuUpdateListingTmallItem struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

}

func (s *TaobaoSkuUpdateListingTmallItem) SetItemId(v int64) *TaobaoSkuUpdateListingTmallItem {
    s.ItemId = &v
    return s
}
