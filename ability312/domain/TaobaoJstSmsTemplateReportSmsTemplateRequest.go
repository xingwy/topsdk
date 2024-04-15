package domain


type TaobaoJstSmsTemplateReportSmsTemplateRequest struct {
    /*
        0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息     */
    TemplateType  *string `json:"template_type,omitempty" `

    /*
        模板名称     */
    TemplateName  *string `json:"template_name,omitempty" `

    /*
        模板内容     */
    TemplateContent  *string `json:"template_content,omitempty" `

    /*
        模板CODE     */
    TemplateCode  *string `json:"template_code,omitempty" `

    /*
        描述信息     */
    Desc  *string `json:"desc,omitempty" `

    /*
        1-- 普通模板  2--数字短信模板     */
    TemplateClass  *string `json:"template_class,omitempty" `

}

func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetTemplateType(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.TemplateType = &v
    return s
}
func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetTemplateName(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.TemplateName = &v
    return s
}
func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetTemplateContent(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.TemplateContent = &v
    return s
}
func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetTemplateCode(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.TemplateCode = &v
    return s
}
func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetDesc(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoJstSmsTemplateReportSmsTemplateRequest) SetTemplateClass(v string) *TaobaoJstSmsTemplateReportSmsTemplateRequest {
    s.TemplateClass = &v
    return s
}
