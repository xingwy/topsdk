package domain


type TaobaoItemcatsGetFeature struct {
    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

}

func (s *TaobaoItemcatsGetFeature) SetAttrKey(v string) *TaobaoItemcatsGetFeature {
    s.AttrKey = &v
    return s
}
func (s *TaobaoItemcatsGetFeature) SetAttrValue(v string) *TaobaoItemcatsGetFeature {
    s.AttrValue = &v
    return s
}
