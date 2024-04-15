package domain


type TaobaoSubuserDutysGetDuty struct {
    /*
        职务ID     */
    DutyId  *int64 `json:"duty_id,omitempty" `

    /*
        职务级别     */
    DutyLevel  *int64 `json:"duty_level,omitempty" `

    /*
        职务名称     */
    DutyName  *string `json:"duty_name,omitempty" `

}

func (s *TaobaoSubuserDutysGetDuty) SetDutyId(v int64) *TaobaoSubuserDutysGetDuty {
    s.DutyId = &v
    return s
}
func (s *TaobaoSubuserDutysGetDuty) SetDutyLevel(v int64) *TaobaoSubuserDutysGetDuty {
    s.DutyLevel = &v
    return s
}
func (s *TaobaoSubuserDutysGetDuty) SetDutyName(v string) *TaobaoSubuserDutysGetDuty {
    s.DutyName = &v
    return s
}
