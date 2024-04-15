package request

import (
	"github.com/xingwy/topsdk/defaultability/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoJstSmsTemplateModifyRequest struct {
	/*
	   修改模板的入参     */
	ModifySmsTemplateRequest *domain.TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest `json:"modify_sms_template_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsTemplateModifyRequest) SetModifySmsTemplateRequest(v domain.TaobaoJstSmsTemplateModifyTopModifySmsTemplateRequest) *TaobaoJstSmsTemplateModifyRequest {
	s.ModifySmsTemplateRequest = &v
	return s
}

func (req *TaobaoJstSmsTemplateModifyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ModifySmsTemplateRequest != nil {
		paramMap["modify_sms_template_request"] = util.ConvertStruct(*req.ModifySmsTemplateRequest)
	}
	return paramMap
}

func (req *TaobaoJstSmsTemplateModifyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
