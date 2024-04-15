package request


type AlibabaAliqinTaVoiceNumDoublecallRequest struct {
    /*
        主叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500     */
    CallerNum  *string `json:"caller_num" required:"true" `
    /*
        主叫号码侧的号码显示，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500     */
    CallerShowNum  *string `json:"caller_show_num" required:"true" `
    /*
        被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500     */
    CalledNum  *string `json:"called_num" required:"true" `
    /*
        被叫号码侧的号码显示，传入的显示号码可以是阿里大于“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。     */
    CalledShowNum  *string `json:"called_show_num" required:"true" `
    /*
        公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。     */
    SessionTimeOut  *string `json:"session_time_out,omitempty" required:"false" `
}

func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetCallerNum(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.CallerNum = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetCallerShowNum(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.CallerShowNum = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetCalledNum(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.CalledNum = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetCalledShowNum(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.CalledShowNum = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetExtend(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.Extend = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallRequest) SetSessionTimeOut(v string) *AlibabaAliqinTaVoiceNumDoublecallRequest {
    s.SessionTimeOut = &v
    return s
}

func (req *AlibabaAliqinTaVoiceNumDoublecallRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CallerNum != nil) {
        paramMap["caller_num"] = *req.CallerNum
    }
    if(req.CallerShowNum != nil) {
        paramMap["caller_show_num"] = *req.CallerShowNum
    }
    if(req.CalledNum != nil) {
        paramMap["called_num"] = *req.CalledNum
    }
    if(req.CalledShowNum != nil) {
        paramMap["called_show_num"] = *req.CalledShowNum
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.SessionTimeOut != nil) {
        paramMap["session_time_out"] = *req.SessionTimeOut
    }
    return paramMap
}

func (req *AlibabaAliqinTaVoiceNumDoublecallRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}