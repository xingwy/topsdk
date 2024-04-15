package domain


type AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO struct {
    /*
        ACCEPT(已揽收),TRANSPORT(运输中),DELIVERING(派送中),SIGN(已签收),CANCEL(已取消),FAILED(物流异常)     */
    Action  *string `json:"action,omitempty" `

    /*
        操作时间戳，精确到毫秒（ms）     */
    OperateTime  *int64 `json:"operate_time,omitempty" `

    /*
        配送员信息     */
    Delivery  *AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO `json:"delivery,omitempty" `

    /*
        货物所在的当前位置     */
    Location  *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO `json:"location,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO) SetAction(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO {
    s.Action = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO) SetOperateTime(v int64) *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO {
    s.OperateTime = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO) SetDelivery(v AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO) *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO {
    s.Delivery = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO) SetLocation(v AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) *AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO {
    s.Location = &v
    return s
}
