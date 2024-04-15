package domain


type AlibabaAscpLogisticsOfflineSendConsignStatusRequest struct {
    /*
        子订单id（组合品不需要传系统会自动计算）     */
    SubTid  *string `json:"sub_tid,omitempty" `

    /*
        子订单是否部分发货，true：部分发货；false：全部发货；周期购、分销订单不支持部分发货     */
    IsPartConsign  *bool `json:"is_part_consign,omitempty" `

}

func (s *AlibabaAscpLogisticsOfflineSendConsignStatusRequest) SetSubTid(v string) *AlibabaAscpLogisticsOfflineSendConsignStatusRequest {
    s.SubTid = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendConsignStatusRequest) SetIsPartConsign(v bool) *AlibabaAscpLogisticsOfflineSendConsignStatusRequest {
    s.IsPartConsign = &v
    return s
}
