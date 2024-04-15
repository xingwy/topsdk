package domain


type AlibabaAscpLogisticsSellerWriteoffResultDTO struct {
    /*
        是否核销成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerWriteoffResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsSellerWriteoffResultDTO {
    s.Success = &v
    return s
}
