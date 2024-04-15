package domain


type AlibabaAscpLogisticsCpGetLogisticsResourceDTO struct {
    /*
        运单号校验正则表达式     */
    RegMailNo  *string `json:"reg_mail_no,omitempty" `

    /*
        快递资源编码     */
    ResourceCode  *string `json:"resource_code,omitempty" `

    /*
        快递资源名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

    /*
        快递公司id     */
    CompanyId  *int64 `json:"company_id,omitempty" `

}

func (s *AlibabaAscpLogisticsCpGetLogisticsResourceDTO) SetRegMailNo(v string) *AlibabaAscpLogisticsCpGetLogisticsResourceDTO {
    s.RegMailNo = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetLogisticsResourceDTO) SetResourceCode(v string) *AlibabaAscpLogisticsCpGetLogisticsResourceDTO {
    s.ResourceCode = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetLogisticsResourceDTO) SetResourceName(v string) *AlibabaAscpLogisticsCpGetLogisticsResourceDTO {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetLogisticsResourceDTO) SetCompanyId(v int64) *AlibabaAscpLogisticsCpGetLogisticsResourceDTO {
    s.CompanyId = &v
    return s
}
