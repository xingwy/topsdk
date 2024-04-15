package request

import (
	"github.com/xingwy/topsdk/defaultability/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoJstSmsSignnameCreateRequest struct {
	/*
	   创建签名入参     */
	AddSmsSignRequest *domain.TaobaoJstSmsSignnameCreateTopAddSmsSignRequest `json:"add_sms_sign_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsSignnameCreateRequest) SetAddSmsSignRequest(v domain.TaobaoJstSmsSignnameCreateTopAddSmsSignRequest) *TaobaoJstSmsSignnameCreateRequest {
	s.AddSmsSignRequest = &v
	return s
}

func (req *TaobaoJstSmsSignnameCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AddSmsSignRequest != nil {
		paramMap["add_sms_sign_request"] = util.ConvertStruct(*req.AddSmsSignRequest)
	}
	return paramMap
}

func (req *TaobaoJstSmsSignnameCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
