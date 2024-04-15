package domain


type AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest struct {
    /*
        运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入     */
    OutSid  *string `json:"out_sid,omitempty" `

    /*
        物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取     */
    CompanyCode  *string `json:"company_code,omitempty" `

    /*
        包裹中商品信息     */
    Goods  *[]AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest `json:"goods,omitempty" `

    /*
        快递子单运单号     */
    SubOutSid  *string `json:"sub_out_sid,omitempty" `

}

func (s *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest) SetOutSid(v string) *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest {
    s.OutSid = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest) SetCompanyCode(v string) *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest {
    s.CompanyCode = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest) SetGoods(v []AlibabaAscpLogisticsOfflineSendTopConsignGoodsRequest) *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest {
    s.Goods = &v
    return s
}
func (s *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest) SetSubOutSid(v string) *AlibabaAscpLogisticsOfflineSendTopConsignPkgRequest {
    s.SubOutSid = &v
    return s
}
