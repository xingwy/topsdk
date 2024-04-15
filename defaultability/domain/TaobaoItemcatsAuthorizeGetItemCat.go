package domain


type TaobaoItemcatsAuthorizeGetItemCat struct {
    /*
        商品所属类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        父类目ID=0时，代表的是一级的类目     */
    ParentCid  *int64 `json:"parent_cid,omitempty" `

    /*
        类目名称     */
    Name  *string `json:"name,omitempty" `

    /*
        状态。可选值:normal(正常),deleted(删除)     */
    Status  *string `json:"status,omitempty" `

    /*
        排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数     */
    SortOrder  *int64 `json:"sort_order,omitempty" `

    /*
        该类目是否为父类目(即：该类目是否还有子类目)     */
    IsParent  *bool `json:"is_parent,omitempty" `

}

func (s *TaobaoItemcatsAuthorizeGetItemCat) SetCid(v int64) *TaobaoItemcatsAuthorizeGetItemCat {
    s.Cid = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetItemCat) SetParentCid(v int64) *TaobaoItemcatsAuthorizeGetItemCat {
    s.ParentCid = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetItemCat) SetName(v string) *TaobaoItemcatsAuthorizeGetItemCat {
    s.Name = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetItemCat) SetStatus(v string) *TaobaoItemcatsAuthorizeGetItemCat {
    s.Status = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetItemCat) SetSortOrder(v int64) *TaobaoItemcatsAuthorizeGetItemCat {
    s.SortOrder = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetItemCat) SetIsParent(v bool) *TaobaoItemcatsAuthorizeGetItemCat {
    s.IsParent = &v
    return s
}
