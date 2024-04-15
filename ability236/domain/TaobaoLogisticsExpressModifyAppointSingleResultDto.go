package domain


type TaobaoLogisticsExpressModifyAppointSingleResultDto struct {
    /*
        是否需要重试     */
    IsRetry  *bool `json:"is_retry,omitempty" `

    /*
        业务返回结果     */
    Result  *TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto `json:"result,omitempty" `

    /*
        错误描述     */
    ErrorDesc  *string `json:"error_desc,omitempty" `

    /*
        调用码     */
    TraceId  *string `json:"trace_id,omitempty" `

    /*
        错误编码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否幂等     */
    IsIdempotent  *bool `json:"is_idempotent,omitempty" `

    /*
        是否调用成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetIsRetry(v bool) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.IsRetry = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetResult(v TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopResponseDto) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.Result = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetErrorDesc(v string) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.ErrorDesc = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetTraceId(v string) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.TraceId = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetErrorCode(v string) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetIsIdempotent(v bool) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.IsIdempotent = &v
    return s
}
func (s *TaobaoLogisticsExpressModifyAppointSingleResultDto) SetSuccess(v bool) *TaobaoLogisticsExpressModifyAppointSingleResultDto {
    s.Success = &v
    return s
}
