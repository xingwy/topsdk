package domain


type TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest struct {
    /*
        不能修改     */
    TemplateType  *int64 `json:"template_type,omitempty" `

    /*
        使用场景说明，可以修改     */
    Remark  *string `json:"remark,omitempty" `

    /*
        不能修改     */
    TemplateCode  *string `json:"template_code,omitempty" `

    /*
        模板名称，可以修改     */
    TemplateName  *string `json:"template_name,omitempty" `

    /*
        模板内容，占位符用${}标识，可以修改     */
    TemplateContent  *string `json:"template_content,omitempty" `

}

func (s *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) SetTemplateType(v int64) *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest {
    s.TemplateType = &v
    return s
}
func (s *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) SetRemark(v string) *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) SetTemplateCode(v string) *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest {
    s.TemplateCode = &v
    return s
}
func (s *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) SetTemplateName(v string) *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest {
    s.TemplateName = &v
    return s
}
func (s *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) SetTemplateContent(v string) *TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest {
    s.TemplateContent = &v
    return s
}
