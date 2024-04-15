package domain


type TaobaoTmcMessagesProduceTmcProduceResult struct {
    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        是否成功     */
    IsSuccess  *bool `json:"is_success,omitempty" `

}

func (s *TaobaoTmcMessagesProduceTmcProduceResult) SetErrorCode(v string) *TaobaoTmcMessagesProduceTmcProduceResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcProduceResult) SetErrorMessage(v string) *TaobaoTmcMessagesProduceTmcProduceResult {
    s.ErrorMessage = &v
    return s
}
func (s *TaobaoTmcMessagesProduceTmcProduceResult) SetIsSuccess(v bool) *TaobaoTmcMessagesProduceTmcProduceResult {
    s.IsSuccess = &v
    return s
}
