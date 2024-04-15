package domain


type TaobaoQimenEventsProduceQimenResult struct {
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

func (s *TaobaoQimenEventsProduceQimenResult) SetErrorCode(v string) *TaobaoQimenEventsProduceQimenResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoQimenEventsProduceQimenResult) SetErrorMessage(v string) *TaobaoQimenEventsProduceQimenResult {
    s.ErrorMessage = &v
    return s
}
func (s *TaobaoQimenEventsProduceQimenResult) SetIsSuccess(v bool) *TaobaoQimenEventsProduceQimenResult {
    s.IsSuccess = &v
    return s
}
