package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoJstSmsTemplateDeleteRequest struct {
    /*
        删除模板的入参     */
    DeleteSmsTemplateRequest  *domain.TaobaoJstSmsTemplateDeleteTopDeleteSmsTemplateRequest `json:"delete_sms_template_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsTemplateDeleteRequest) SetDeleteSmsTemplateRequest(v domain.TaobaoJstSmsTemplateDeleteTopDeleteSmsTemplateRequest) *TaobaoJstSmsTemplateDeleteRequest {
    s.DeleteSmsTemplateRequest = &v
    return s
}

func (req *TaobaoJstSmsTemplateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DeleteSmsTemplateRequest != nil) {
        paramMap["delete_sms_template_request"] = util.ConvertStruct(*req.DeleteSmsTemplateRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsTemplateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}