package domain


type AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO struct {
    /*
        配送员电话，支持手机、座机、400电话     */
    Phone  *string `json:"phone,omitempty" `

    /*
        配送员姓名     */
    Name  *string `json:"name,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO) SetPhone(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO {
    s.Phone = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO) SetName(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeDeliveryTopDTO {
    s.Name = &v
    return s
}
