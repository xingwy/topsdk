package request

import (
	"github.com/xingwy/topsdk/defaultability/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoJstSmsSignnameModifyRequest struct {
	/*
	   修改签名入参     */
	ModifySmsSignRequest *domain.TaobaoJstSmsSignnameModifyTopModifySmsSignRequest `json:"modify_sms_sign_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsSignnameModifyRequest) SetModifySmsSignRequest(v domain.TaobaoJstSmsSignnameModifyTopModifySmsSignRequest) *TaobaoJstSmsSignnameModifyRequest {
	s.ModifySmsSignRequest = &v
	return s
}

func (req *TaobaoJstSmsSignnameModifyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ModifySmsSignRequest != nil {
		paramMap["modify_sms_sign_request"] = util.ConvertStruct(*req.ModifySmsSignRequest)
	}
	return paramMap
}

func (req *TaobaoJstSmsSignnameModifyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
