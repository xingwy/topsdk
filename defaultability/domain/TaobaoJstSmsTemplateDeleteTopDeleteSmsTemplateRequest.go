package domain


type TaobaoJstSmsTemplateDeleteTopDeleteSmsTemplateRequest struct {
    /*
        待删除的模板code     */
    TemplateCode  *string `json:"template_code,omitempty" `

}

func (s *TaobaoJstSmsTemplateDeleteTopDeleteSmsTemplateRequest) SetTemplateCode(v string) *TaobaoJstSmsTemplateDeleteTopDeleteSmsTemplateRequest {
    s.TemplateCode = &v
    return s
}
