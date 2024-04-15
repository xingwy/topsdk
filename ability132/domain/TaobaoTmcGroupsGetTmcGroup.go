package domain


type TaobaoTmcGroupsGetTmcGroup struct {
    /*
        分组名称     */
    Name  *string `json:"name,omitempty" `

}

func (s *TaobaoTmcGroupsGetTmcGroup) SetName(v string) *TaobaoTmcGroupsGetTmcGroup {
    s.Name = &v
    return s
}
