package request

import (
	"github.com/xingwy/topsdk/ability312/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoJstSmsTemplateReportRequest struct {
	/*
	   存量短信模板上报入参     */
	SmsTemplateRequest *[]domain.TaobaoJstSmsTemplateReportSmsTemplateRequest `json:"sms_template_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsTemplateReportRequest) SetSmsTemplateRequest(v []domain.TaobaoJstSmsTemplateReportSmsTemplateRequest) *TaobaoJstSmsTemplateReportRequest {
	s.SmsTemplateRequest = &v
	return s
}

func (req *TaobaoJstSmsTemplateReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SmsTemplateRequest != nil {
		paramMap["sms_template_request"] = util.ConvertStructList(*req.SmsTemplateRequest)
	}
	return paramMap
}

func (req *TaobaoJstSmsTemplateReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
