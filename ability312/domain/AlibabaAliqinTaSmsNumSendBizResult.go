package domain


type AlibabaAliqinTaSmsNumSendBizResult struct {
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

func (s *AlibabaAliqinTaSmsNumSendBizResult) SetModel(v string) *AlibabaAliqinTaSmsNumSendBizResult {
    s.Model = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendBizResult) SetMsg(v string) *AlibabaAliqinTaSmsNumSendBizResult {
    s.Msg = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendBizResult) SetSuccess(v bool) *AlibabaAliqinTaSmsNumSendBizResult {
    s.Success = &v
    return s
}
func (s *AlibabaAliqinTaSmsNumSendBizResult) SetErrCode(v string) *AlibabaAliqinTaSmsNumSendBizResult {
    s.ErrCode = &v
    return s
}
