package domain


type TaobaoJstSmsTemplateQueryAccessBaseDTO struct {
    /*
        0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息     */
    TemplateType  *int64 `json:"template_type,omitempty" `

    /*
        0--待审核  1--通过  2--拒绝     */
    TemplateStatus  *int64 `json:"template_status,omitempty" `

    /*
        审核意见     */
    Reason  *string `json:"reason,omitempty" `

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
        创建时间     */
    CreateDate  *string `json:"create_date,omitempty" `

}

func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetTemplateType(v int64) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.TemplateType = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetTemplateStatus(v int64) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.TemplateStatus = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetReason(v string) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.Reason = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetTemplateName(v string) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.TemplateName = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetTemplateContent(v string) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.TemplateContent = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetTemplateCode(v string) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.TemplateCode = &v
    return s
}
func (s *TaobaoJstSmsTemplateQueryAccessBaseDTO) SetCreateDate(v string) *TaobaoJstSmsTemplateQueryAccessBaseDTO {
    s.CreateDate = &v
    return s
}
