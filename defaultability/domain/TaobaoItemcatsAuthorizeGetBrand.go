package domain


type TaobaoItemcatsAuthorizeGetBrand struct {
    /*
        对应属性的 pid:vid 串中的vid     */
    Vid  *int64 `json:"vid,omitempty" `

    /*
        vid的值     */
    Name  *string `json:"name,omitempty" `

    /*
        品牌的属性id     */
    Pid  *int64 `json:"pid,omitempty" `

    /*
        品牌的属性名     */
    PropName  *string `json:"prop_name,omitempty" `

}

func (s *TaobaoItemcatsAuthorizeGetBrand) SetVid(v int64) *TaobaoItemcatsAuthorizeGetBrand {
    s.Vid = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetBrand) SetName(v string) *TaobaoItemcatsAuthorizeGetBrand {
    s.Name = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetBrand) SetPid(v int64) *TaobaoItemcatsAuthorizeGetBrand {
    s.Pid = &v
    return s
}
func (s *TaobaoItemcatsAuthorizeGetBrand) SetPropName(v string) *TaobaoItemcatsAuthorizeGetBrand {
    s.PropName = &v
    return s
}
