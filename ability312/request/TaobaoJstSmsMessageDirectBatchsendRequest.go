package request


type TaobaoJstSmsMessageDirectBatchsendRequest struct {
    /*
        短信签名（如果是通过OAID发送短信则签名需要是数组格式，数组长度需要和OAID数量保持一致。）     */
    SignName  *string `json:"sign_name" required:"true" `
    /*
        废弃字段     */
    Url  *string `json:"url,omitempty" required:"false" `
    /*
        短信模版Code（明文发送短信和OAID发送均不要传数组格式）     */
    SmsTemplateCode  *string `json:"sms_template_code" required:"true" `
    /*
        短信接收号码，json格式，最多200个号码     */
    RecNum  *string `json:"rec_num,omitempty" required:"false" `
    /*
        模板参数替换方式："[{\"msg\":\"hello1\",\"date\":\"2021-12-03\"},{\"msg\":\"hello2\",\"date\":\"2021-12-04\"},{\"msg\":\"hello3\",\"date\":\"2021-12-05\"}]"     */
    SmsContent  *string `json:"sms_content,omitempty" required:"false" `
    /*
        短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。     */
    ExtendNum  *string `json:"extend_num,omitempty" required:"false" `
    /*
        废弃字段     */
    TaskCode  *string `json:"task_code,omitempty" required:"false" `
    /*
        对于taskSign相同的请求平台认为是商家的同一次短信发送，可用于对OAID的明文号码去重。     */
    TaskSign  *string `json:"task_sign,omitempty" required:"false" `
    /*
        OAID批量发短信的入参，传该参数的时候rec_num不需要传，最大50个。     */
    Oaids  *string `json:"oaids,omitempty" required:"false" `
    /*
        OAID批量发短信时必传，为OAID一一对应的订单ID。     */
    OrderIds  *string `json:"order_ids,omitempty" required:"false" `
    /*
        如果传，必须是一个JSON对象格式的字符串。     */
    ExtraData  *string `json:"extra_data,omitempty" required:"false" `
}

func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetSignName(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.SignName = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetUrl(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.Url = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsTemplateCode(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.SmsTemplateCode = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetRecNum(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.RecNum = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsContent(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.SmsContent = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetExtendNum(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.ExtendNum = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskCode(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.TaskCode = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskSign(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.TaskSign = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetOaids(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.Oaids = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetOrderIds(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.OrderIds = &v
    return s
}
func (s *TaobaoJstSmsMessageDirectBatchsendRequest) SetExtraData(v string) *TaobaoJstSmsMessageDirectBatchsendRequest {
    s.ExtraData = &v
    return s
}

func (req *TaobaoJstSmsMessageDirectBatchsendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SignName != nil) {
        paramMap["sign_name"] = *req.SignName
    }
    if(req.Url != nil) {
        paramMap["url"] = *req.Url
    }
    if(req.SmsTemplateCode != nil) {
        paramMap["sms_template_code"] = *req.SmsTemplateCode
    }
    if(req.RecNum != nil) {
        paramMap["rec_num"] = *req.RecNum
    }
    if(req.SmsContent != nil) {
        paramMap["sms_content"] = *req.SmsContent
    }
    if(req.ExtendNum != nil) {
        paramMap["extend_num"] = *req.ExtendNum
    }
    if(req.TaskCode != nil) {
        paramMap["task_code"] = *req.TaskCode
    }
    if(req.TaskSign != nil) {
        paramMap["task_sign"] = *req.TaskSign
    }
    if(req.Oaids != nil) {
        paramMap["oaids"] = *req.Oaids
    }
    if(req.OrderIds != nil) {
        paramMap["order_ids"] = *req.OrderIds
    }
    if(req.ExtraData != nil) {
        paramMap["extra_data"] = *req.ExtraData
    }
    return paramMap
}

func (req *TaobaoJstSmsMessageDirectBatchsendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}