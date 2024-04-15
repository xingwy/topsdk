package domain


type AlibabaAscpLogisticsIdentcodeQueryResultDTO struct {
    /*
        识别码列表     */
    IdentCodeList  *[]AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO `json:"ident_code_list,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAscpLogisticsIdentcodeQueryResultDTO) SetIdentCodeList(v []AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) *AlibabaAscpLogisticsIdentcodeQueryResultDTO {
    s.IdentCodeList = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryResultDTO) SetSuccess(v bool) *AlibabaAscpLogisticsIdentcodeQueryResultDTO {
    s.Success = &v
    return s
}
