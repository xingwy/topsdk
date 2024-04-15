package domain


type AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO struct {
    /*
        识别码     */
    IdentCode  *string `json:"ident_code,omitempty" `

    /*
        识别码类型，SN/IMEI     */
    IdentType  *string `json:"ident_type,omitempty" `

    /*
        品牌ID     */
    BrandId  *string `json:"brand_id,omitempty" `

    /*
        根类目ID     */
    RootCatId  *string `json:"root_cat_id,omitempty" `

    /*
        是否可用     */
    Available  *bool `json:"available,omitempty" `

    /*
        不可用原因     */
    UnAvailableReason  *string `json:"un_available_reason,omitempty" `

}

func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetIdentCode(v string) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.IdentCode = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetIdentType(v string) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.IdentType = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetBrandId(v string) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.BrandId = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetRootCatId(v string) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.RootCatId = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetAvailable(v bool) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.Available = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO) SetUnAvailableReason(v string) *AlibabaAscpLogisticsIdentcodeQueryTopIdentCodeDTO {
    s.UnAvailableReason = &v
    return s
}
