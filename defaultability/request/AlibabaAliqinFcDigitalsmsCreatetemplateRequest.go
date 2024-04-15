package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type AlibabaAliqinFcDigitalsmsCreatetemplateRequest struct {
    /*
        模板名称     */
    TemplateName  *string `json:"template_name" required:"true" `
    /*
        系统自动生成     */
    TemplateContents  *[]domain.AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto `json:"template_contents" required:"true" `
    /*
        申请说明     */
    ApplyRemark  *string `json:"apply_remark" required:"true" `
}

func (s *AlibabaAliqinFcDigitalsmsCreatetemplateRequest) SetTemplateName(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateRequest {
    s.TemplateName = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateRequest) SetTemplateContents(v []domain.AlibabaAliqinFcDigitalsmsCreatetemplateDigitalSmsTemplateContentDto) *AlibabaAliqinFcDigitalsmsCreatetemplateRequest {
    s.TemplateContents = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateRequest) SetApplyRemark(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateRequest {
    s.ApplyRemark = &v
    return s
}

func (req *AlibabaAliqinFcDigitalsmsCreatetemplateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TemplateName != nil) {
        paramMap["template_name"] = *req.TemplateName
    }
    if(req.TemplateContents != nil) {
        paramMap["template_contents"] = util.ConvertStructList(*req.TemplateContents)
    }
    if(req.ApplyRemark != nil) {
        paramMap["apply_remark"] = *req.ApplyRemark
    }
    return paramMap
}

func (req *AlibabaAliqinFcDigitalsmsCreatetemplateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}