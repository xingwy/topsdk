package domain


type TaobaoJstSmsSignnameReportSmsSignNameRequest struct {
    /*
        短信签名     */
    SignName  *string `json:"sign_name,omitempty" `

    /*
        描述信息     */
    Description  *string `json:"description,omitempty" `

}

func (s *TaobaoJstSmsSignnameReportSmsSignNameRequest) SetSignName(v string) *TaobaoJstSmsSignnameReportSmsSignNameRequest {
    s.SignName = &v
    return s
}
func (s *TaobaoJstSmsSignnameReportSmsSignNameRequest) SetDescription(v string) *TaobaoJstSmsSignnameReportSmsSignNameRequest {
    s.Description = &v
    return s
}
