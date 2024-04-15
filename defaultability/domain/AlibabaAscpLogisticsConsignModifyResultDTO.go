package domain


type AlibabaAscpLogisticsConsignModifyResultDTO struct {
    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsConsignModifyResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsConsignModifyResultDTO {
    s.Success = &v
    return s
}
