package request

import (
        "topsdk/util"
    )

type AlibabaAliqinTaSmsNumSendRequest struct {
    /*
        接收号码     */
    RecNum  *string `json:"rec_num" required:"true" `
    /*
        短信模板CODE     */
    SmsTemplateCode  *string `json:"sms_template_code" required:"true" `
    /*
        公共回传参数     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        短信签名     */
    SmsFreeSignName  *string `json:"sms_free_sign_name" required:"true" `
    /*
        短信模板变量，AckNum是变量参数     */
    SmsParam  *map[string]interface{} `json:"sms_param,omitempty" required:"false" `
    /*
        类型，normal：短信     */
    SmsType  *string `json:"sms_type" required:"true" `
    /*
        商家自定义扩展码     */
    ExtendCode  *string `json:"extend_code,omitempty" required:"false" `
    /*
        商家自定义扩展名,例如店铺nick     */
    ExtendName  *string `json:"extend_name,omitempty" required:"false" `
}

func (s *AlibabaAliqinTaSmsNumSendRequest) SetRecNum(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.RecNum = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetSmsTemplateCode(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.SmsTemplateCode = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetExtend(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.Extend = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetSmsFreeSignName(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.SmsFreeSignName = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetSmsParam(v map[string]interface{}) *AlibabaAliqinTaSmsNumSendRequest {
    s.SmsParam = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetSmsType(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.SmsType = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetExtendCode(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.ExtendCode = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendRequest) SetExtendName(v string) *AlibabaAliqinTaSmsNumSendRequest {
    s.ExtendName = &v
    return s
}

func (req *AlibabaAliqinTaSmsNumSendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RecNum != nil) {
        paramMap["rec_num"] = *req.RecNum
    }
    if(req.SmsTemplateCode != nil) {
        paramMap["sms_template_code"] = *req.SmsTemplateCode
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.SmsFreeSignName != nil) {
        paramMap["sms_free_sign_name"] = *req.SmsFreeSignName
    }
    if(req.SmsParam != nil) {
        paramMap["sms_param"] = util.ConvertStruct(*req.SmsParam)
    }
    if(req.SmsType != nil) {
        paramMap["sms_type"] = *req.SmsType
    }
    if(req.ExtendCode != nil) {
        paramMap["extend_code"] = *req.ExtendCode
    }
    if(req.ExtendName != nil) {
        paramMap["extend_name"] = *req.ExtendName
    }
    return paramMap
}

func (req *AlibabaAliqinTaSmsNumSendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}