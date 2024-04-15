package domain


type TaobaoJstSmsSignnameQueryTopQuerySmsSignRequest struct {
    /*
        要查询的签名     */
    SignName  *string `json:"sign_name,omitempty" `

}

func (s *TaobaoJstSmsSignnameQueryTopQuerySmsSignRequest) SetSignName(v string) *TaobaoJstSmsSignnameQueryTopQuerySmsSignRequest {
    s.SignName = &v
    return s
}
