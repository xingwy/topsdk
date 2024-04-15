package domain


type TaobaoJstSmsSignnameDeleteTopDeleteSmsSignRequest struct {
    /*
        待删除的签名     */
    SignName  *string `json:"sign_name,omitempty" `

}

func (s *TaobaoJstSmsSignnameDeleteTopDeleteSmsSignRequest) SetSignName(v string) *TaobaoJstSmsSignnameDeleteTopDeleteSmsSignRequest {
    s.SignName = &v
    return s
}
