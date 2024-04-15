package domain


type TaobaoJstSmsSignnameQueryQuerySmsSignDTO struct {
    /*
        拒绝     */
    SignStatus  *int64 `json:"sign_status,omitempty" `

    /*
        被拒绝的原因     */
    Reason  *string `json:"reason,omitempty" `

    /*
        申请的签名     */
    SignName  *string `json:"sign_name,omitempty" `

    /*
        签名创建时间     */
    CreateDate  *string `json:"create_date,omitempty" `

}

func (s *TaobaoJstSmsSignnameQueryQuerySmsSignDTO) SetSignStatus(v int64) *TaobaoJstSmsSignnameQueryQuerySmsSignDTO {
    s.SignStatus = &v
    return s
}
func (s *TaobaoJstSmsSignnameQueryQuerySmsSignDTO) SetReason(v string) *TaobaoJstSmsSignnameQueryQuerySmsSignDTO {
    s.Reason = &v
    return s
}
func (s *TaobaoJstSmsSignnameQueryQuerySmsSignDTO) SetSignName(v string) *TaobaoJstSmsSignnameQueryQuerySmsSignDTO {
    s.SignName = &v
    return s
}
func (s *TaobaoJstSmsSignnameQueryQuerySmsSignDTO) SetCreateDate(v string) *TaobaoJstSmsSignnameQueryQuerySmsSignDTO {
    s.CreateDate = &v
    return s
}
