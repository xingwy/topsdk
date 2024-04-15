package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoJstSmsTemplateCreateRequest struct {
    /*
        申请模板入参     */
    SmsTemplateForIsvRequest  *domain.TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest `json:"sms_template_for_isv_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsTemplateCreateRequest) SetSmsTemplateForIsvRequest(v domain.TaobaoJstSmsTemplateCreateAddSmsTemplateForIsvRequest) *TaobaoJstSmsTemplateCreateRequest {
    s.SmsTemplateForIsvRequest = &v
    return s
}

func (req *TaobaoJstSmsTemplateCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SmsTemplateForIsvRequest != nil) {
        paramMap["sms_template_for_isv_request"] = util.ConvertStruct(*req.SmsTemplateForIsvRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsTemplateCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}