package domain


type AlibabaAliqinTaNumberSinglecallbyvoiceResult struct {
    /*
        结果     */
    Model  *string `json:"model,omitempty" `

    /*
        系统自动生成     */
    Msg  *string `json:"msg,omitempty" `

    /*
        成功，失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        系统自动生成     */
    Code  *string `json:"code,omitempty" `

}

func (s *AlibabaAliqinTaNumberSinglecallbyvoiceResult) SetModel(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceResult {
    s.Model = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceResult) SetMsg(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceResult {
    s.Msg = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceResult) SetSuccess(v bool) *AlibabaAliqinTaNumberSinglecallbyvoiceResult {
    s.Success = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceResult) SetCode(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceResult {
    s.Code = &v
    return s
}
