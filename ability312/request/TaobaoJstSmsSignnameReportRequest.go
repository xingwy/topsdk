package request

import (
        "topsdk/ability312/domain"
        "topsdk/util"
    )

type TaobaoJstSmsSignnameReportRequest struct {
    /*
        上传签名入参     */
    SmsSignNameRequest  *[]domain.TaobaoJstSmsSignnameReportSmsSignNameRequest `json:"sms_sign_name_request,omitempty" required:"false" `
}

func (s *TaobaoJstSmsSignnameReportRequest) SetSmsSignNameRequest(v []domain.TaobaoJstSmsSignnameReportSmsSignNameRequest) *TaobaoJstSmsSignnameReportRequest {
    s.SmsSignNameRequest = &v
    return s
}

func (req *TaobaoJstSmsSignnameReportRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SmsSignNameRequest != nil) {
        paramMap["sms_sign_name_request"] = util.ConvertStructList(*req.SmsSignNameRequest)
    }
    return paramMap
}

func (req *TaobaoJstSmsSignnameReportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}