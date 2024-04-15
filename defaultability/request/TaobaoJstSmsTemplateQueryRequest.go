package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoJstSmsTemplateQueryRequest struct {
    /*
        查询短信模板的入参     */
    QuerySmsTemplateRequest  *domain.TaobaoJstSmsTemplateQueryTopQuerySmsTemplateRequest `json:"query_sms_template_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsTemplateQueryRequest) SetQuerySmsTemplateRequest(v domain.TaobaoJstSmsTemplateQueryTopQuerySmsTemplateRequest) *TaobaoJstSmsTemplateQueryRequest {
    s.QuerySmsTemplateRequest = &v
    return s
}

func (req *TaobaoJstSmsTemplateQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.QuerySmsTemplateRequest != nil) {
        paramMap["query_sms_template_request"] = util.ConvertStruct(*req.QuerySmsTemplateRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsTemplateQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}