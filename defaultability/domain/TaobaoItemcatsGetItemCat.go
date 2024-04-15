package domain


type TaobaoItemcatsGetItemCat struct {
    /*
        商品所属类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        Feature对象列表
目前已有的属性：
若Attr_key为 udsaleprop，attr_value为1 则允许卖家在改类目新增自定义销售属性,不然为不允许     */
    Features  *[]TaobaoItemcatsGetFeature `json:"features,omitempty" `

    /*
        该类目是否为父类目(即：该类目是否还有子类目)     */
    IsParent  *bool `json:"is_parent,omitempty" `

    /*
        类目名称     */
    Name  *string `json:"name,omitempty" `

    /*
        父类目ID=0时，代表的是一级的类目     */
    ParentCid  *int64 `json:"parent_cid,omitempty" `

    /*
        排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数     */
    SortOrder  *int64 `json:"sort_order,omitempty" `

    /*
        状态。可选值:normal(正常),deleted(删除)     */
    Status  *string `json:"status,omitempty" `

    /*
        是否度量衡类目     */
    TaosirCat  *bool `json:"taosir_cat,omitempty" `

}

func (s *TaobaoItemcatsGetItemCat) SetCid(v int64) *TaobaoItemcatsGetItemCat {
    s.Cid = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetFeatures(v []TaobaoItemcatsGetFeature) *TaobaoItemcatsGetItemCat {
    s.Features = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetIsParent(v bool) *TaobaoItemcatsGetItemCat {
    s.IsParent = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetName(v string) *TaobaoItemcatsGetItemCat {
    s.Name = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetParentCid(v int64) *TaobaoItemcatsGetItemCat {
    s.ParentCid = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetSortOrder(v int64) *TaobaoItemcatsGetItemCat {
    s.SortOrder = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetStatus(v string) *TaobaoItemcatsGetItemCat {
    s.Status = &v
    return s
}
func (s *TaobaoItemcatsGetItemCat) SetTaosirCat(v bool) *TaobaoItemcatsGetItemCat {
    s.TaosirCat = &v
    return s
}
