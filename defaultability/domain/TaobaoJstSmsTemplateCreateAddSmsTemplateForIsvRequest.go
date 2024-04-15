package domain


type TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest struct {
    /*
        0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息     */
    TemplateType  *int64 `json:"template_type,omitempty" `

    /*
        NORMAL -- 普通模板  DIGITAL--数字短信模板 defalutValue:NORMAL    */
    TemplateTypeClass  *string `json:"template_type_class,omitempty" `

    /*
        上传文件     */
    TemplateInfos  *[]TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO `json:"template_infos,omitempty" `

    /*
        使用场景说明     */
    Remark  *string `json:"remark,omitempty" `

    /*
        模板名称     */
    TemplateName  *string `json:"template_name,omitempty" `

    /*
        模板内容，占位符用${}标识     */
    TemplateContent  *string `json:"template_content,omitempty" `

}

func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetTemplateType(v int64) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.TemplateType = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetTemplateTypeClass(v string) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.TemplateTypeClass = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetTemplateInfos(v []TaobaoJstSmsTemplateCreateDigitalSmsTemplateContentDTO) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.TemplateInfos = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetRemark(v string) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetTemplateName(v string) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.TemplateName = &v
    return s
}
func (s *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) SetTemplateContent(v string) *TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest {
    s.TemplateContent = &v
    return s
}
