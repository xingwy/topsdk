package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoJstSmsSignnameQueryRequest struct {
    /*
        签名查询的入参     */
    QuerySmsSignRequest  *domain.TaobaoJstSmsSignnameQueryTopQuerySmsSignRequest `json:"query_sms_sign_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsSignnameQueryRequest) SetQuerySmsSignRequest(v domain.TaobaoJstSmsSignnameQueryTopQuerySmsSignRequest) *TaobaoJstSmsSignnameQueryRequest {
    s.QuerySmsSignRequest = &v
    return s
}

func (req *TaobaoJstSmsSignnameQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.QuerySmsSignRequest != nil) {
        paramMap["query_sms_sign_request"] = util.ConvertStruct(*req.QuerySmsSignRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsSignnameQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}