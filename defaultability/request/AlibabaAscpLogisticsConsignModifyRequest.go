package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type AlibabaAscpLogisticsConsignModifyRequest struct {
    /*
        订单id     */
    Tid  *string `json:"tid" required:"true" `
    /*
        原物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取     */
    OldCompanyCode  *string `json:"old_company_code,omitempty" required:"false" `
    /*
        原运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入     */
    OldOutSid  *string `json:"old_out_sid" required:"true" `
    /*
        新物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取     */
    NewCompanyCode  *string `json:"new_company_code,omitempty" required:"false" `
    /*
        新运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入     */
    NewOutSid  *string `json:"new_out_sid" required:"true" `
    /*
        原包裹中的商品信息     */
    Goods  *[]domain.AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest `json:"goods,omitempty" required:"false" `
    /*
        feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号"12345678910-1234"。     */
    Feature  *string `json:"feature,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsConsignModifyRequest) SetTid(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.Tid = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetOldCompanyCode(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.OldCompanyCode = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetOldOutSid(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.OldOutSid = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetNewCompanyCode(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.NewCompanyCode = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetNewOutSid(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.NewOutSid = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetGoods(v []domain.AlibabaAscpLogisticsConsignModifyTopConsignGoodsRequest) *AlibabaAscpLogisticsConsignModifyRequest {
    s.Goods = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignModifyRequest) SetFeature(v string) *AlibabaAscpLogisticsConsignModifyRequest {
    s.Feature = &v
    return s
}

func (req *AlibabaAscpLogisticsConsignModifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.OldCompanyCode != nil) {
        paramMap["old_company_code"] = *req.OldCompanyCode
    }
    if(req.OldOutSid != nil) {
        paramMap["old_out_sid"] = *req.OldOutSid
    }
    if(req.NewCompanyCode != nil) {
        paramMap["new_company_code"] = *req.NewCompanyCode
    }
    if(req.NewOutSid != nil) {
        paramMap["new_out_sid"] = *req.NewOutSid
    }
    if(req.Goods != nil) {
        paramMap["goods"] = util.ConvertStructList(*req.Goods)
    }
    if(req.Feature != nil) {
        paramMap["feature"] = *req.Feature
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsConsignModifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}