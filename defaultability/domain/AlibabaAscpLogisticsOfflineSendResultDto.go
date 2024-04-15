package domain


type AlibabaAscpLogisticsOfflineSendResultDto struct {
    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsOfflineSendResultDto) SetSuccess(v bool) *AlibabaAscpLogisticsOfflineSendResultDto {
    s.Success = &v
    return s
}
