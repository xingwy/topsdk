package domain


type AlibabaAliqinTaVoiceNumDoublecallBizResult struct {
    /*
        返回结果     */
    Model  *string `json:"model,omitempty" `

    /*
        返回信息描述     */
    Msg  *string `json:"msg,omitempty" `

    /*
        true表示成功，false表示失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

}

func (s *AlibabaAliqinTaVoiceNumDoublecallBizResult) SetModel(v string) *AlibabaAliqinTaVoiceNumDoublecallBizResult {
    s.Model = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallBizResult) SetMsg(v string) *AlibabaAliqinTaVoiceNumDoublecallBizResult {
    s.Msg = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallBizResult) SetSuccess(v bool) *AlibabaAliqinTaVoiceNumDoublecallBizResult {
    s.Success = &v
    return s
}
func (s *AlibabaAliqinTaVoiceNumDoublecallBizResult) SetErrCode(v string) *AlibabaAliqinTaVoiceNumDoublecallBizResult {
    s.ErrCode = &v
    return s
}
