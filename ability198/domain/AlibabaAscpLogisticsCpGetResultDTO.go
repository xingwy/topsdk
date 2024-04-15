package domain


type AlibabaAscpLogisticsCpGetResultDTO struct {
    /*
        快递公司资源列表     */
    ResourceList  *[]AlibabaAscpLogisticsCpGetLogisticsResourceDTO `json:"resource_list,omitempty" `

    /*
        调用是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误信息     */
    ErrorDesc  *string `json:"error_desc,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *AlibabaAscpLogisticsCpGetResultDTO) SetResourceList(v []AlibabaAscpLogisticsCpGetLogisticsResourceDTO) *AlibabaAscpLogisticsCpGetResultDTO {
    s.ResourceList = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsCpGetResultDTO {
    s.Success = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetResultDTO) SetErrorDesc(v string) *AlibabaAscpLogisticsCpGetResultDTO {
    s.ErrorDesc = &v
    return s
}
func (s *AlibabaAscpLogisticsCpGetResultDTO) SetErrorCode(v string) *AlibabaAscpLogisticsCpGetResultDTO {
    s.ErrorCode = &v
    return s
}
