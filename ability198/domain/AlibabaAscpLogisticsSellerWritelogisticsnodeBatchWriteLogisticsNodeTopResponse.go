package domain


type AlibabaAscpLogisticsSellerWritelogisticsnodeBatchWriteLogisticsNodeTopResponse struct {
    /*
        true成功，false失败     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeBatchWriteLogisticsNodeTopResponse) SetSuccess(v bool) *AlibabaAscpLogisticsSellerWritelogisticsnodeBatchWriteLogisticsNodeTopResponse {
    s.Success = &v
    return s
}
