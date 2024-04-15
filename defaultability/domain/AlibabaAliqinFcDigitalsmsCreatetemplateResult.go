package domain


type AlibabaAliqinFcDigitalsmsCreatetemplateResult struct {
    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        true表示成功，false表示失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回信息描述     */
    Msg  *string `json:"msg,omitempty" `

    /*
        模板code     */
    Model  *string `json:"model,omitempty" `

}

func (s *AlibabaAliqinFcDigitalsmsCreatetemplateResult) SetErrCode(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateResult) SetSuccess(v bool) *AlibabaAliqinFcDigitalsmsCreatetemplateResult {
    s.Success = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateResult) SetMsg(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateResult {
    s.Msg = &v
    return s
}
func (s *AlibabaAliqinFcDigitalsmsCreatetemplateResult) SetModel(v string) *AlibabaAliqinFcDigitalsmsCreatetemplateResult {
    s.Model = &v
    return s
}
