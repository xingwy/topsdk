package domain


type AlibabaAliqinTaNumberSinglecallbyttsResult struct {
    /*
        返回值     */
    Model  *string `json:"model,omitempty" `

    /*
        信息     */
    Msg  *string `json:"msg,omitempty" `

    /*
        成功，失败     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAliqinTaNumberSinglecallbyttsResult) SetModel(v string) *AlibabaAliqinTaNumberSinglecallbyttsResult {
    s.Model = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyttsResult) SetMsg(v string) *AlibabaAliqinTaNumberSinglecallbyttsResult {
    s.Msg = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyttsResult) SetSuccess(v bool) *AlibabaAliqinTaNumberSinglecallbyttsResult {
    s.Success = &v
    return s
}
