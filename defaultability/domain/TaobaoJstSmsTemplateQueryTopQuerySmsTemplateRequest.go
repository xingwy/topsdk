package domain


type TaobaoJstSmsTemplateQueryTopQuerySmsTemplateRequest struct {
    /*
        要查询的模板CODE     */
    TemplateCode  *string `json:"template_code,omitempty" `

}

func (s *TaobaoJstSmsTemplateQueryTopQuerySmsTemplateRequest) SetTemplateCode(v string) *TaobaoJstSmsTemplateQueryTopQuerySmsTemplateRequest {
    s.TemplateCode = &v
    return s
}
