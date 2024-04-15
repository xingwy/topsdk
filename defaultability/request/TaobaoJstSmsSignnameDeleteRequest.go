package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoJstSmsSignnameDeleteRequest struct {
    /*
        删除签名入参     */
    DeleteSmsSignRequest  *domain.TaobaoJstSmsSignnameDeleteTopDeleteSmsSignRequest `json:"delete_sms_sign_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsSignnameDeleteRequest) SetDeleteSmsSignRequest(v domain.TaobaoJstSmsSignnameDeleteTopDeleteSmsSignRequest) *TaobaoJstSmsSignnameDeleteRequest {
    s.DeleteSmsSignRequest = &v
    return s
}

func (req *TaobaoJstSmsSignnameDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DeleteSmsSignRequest != nil) {
        paramMap["delete_sms_sign_request"] = util.ConvertStruct(*req.DeleteSmsSignRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsSignnameDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}