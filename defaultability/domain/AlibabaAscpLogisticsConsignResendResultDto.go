package domain


type AlibabaAscpLogisticsConsignResendResultDto struct {
    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsConsignResendResultDto) SetSuccess(v bool) *AlibabaAscpLogisticsConsignResendResultDto {
    s.Success = &v
    return s
}
