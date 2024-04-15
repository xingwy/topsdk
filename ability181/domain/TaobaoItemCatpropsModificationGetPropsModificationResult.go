package domain


type TaobaoItemCatpropsModificationGetPropsModificationResult struct {
    /*
        类目属性Id     */
    PropertyId  *int64 `json:"property_id,omitempty" `

    /*
        是否必填，0-否，1-是     */
    Required  *int64 `json:"required,omitempty" `

    /*
        是否多选，0-否，1-是     */
    MultiSelect  *int64 `json:"multi_select,omitempty" `

    /*
        变更类型: 删除(1), 修改(2), 新增(3)     */
    Type  *int64 `json:"type,omitempty" `

    /*
        变更日期     */
    Ds  *string `json:"ds,omitempty" `

    /*
        属性名称     */
    PropertyName  *string `json:"property_name,omitempty" `

}

func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetPropertyId(v int64) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.PropertyId = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetRequired(v int64) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.Required = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetMultiSelect(v int64) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.MultiSelect = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetType(v int64) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.Type = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetDs(v string) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.Ds = &v
    return s
}
func (s *TaobaoItemCatpropsModificationGetPropsModificationResult) SetPropertyName(v string) *TaobaoItemCatpropsModificationGetPropsModificationResult {
    s.PropertyName = &v
    return s
}
