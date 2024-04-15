package domain


type AlibabaAscpLogisticsSellerSendResultDTO struct {
    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

    /*
        -     */
    Consign  *AlibabaAscpLogisticsSellerSendConsignDTO `json:"consign,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerSendResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsSellerSendResultDTO {
    s.Success = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendResultDTO) SetConsign(v AlibabaAscpLogisticsSellerSendConsignDTO) *AlibabaAscpLogisticsSellerSendResultDTO {
    s.Consign = &v
    return s
}
