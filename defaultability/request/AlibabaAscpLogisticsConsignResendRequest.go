package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type AlibabaAscpLogisticsConsignResendRequest struct {
    /*
        订单id     */
    Tid  *string `json:"tid" required:"true" `
    /*
        拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集     */
    SubTids  *string `json:"sub_tids,omitempty" required:"false" `
    /*
        包裹包含的运单号及快递公司编号,多包裹时，需要包含所有包裹的运单号等信息     */
    ConsignPkgs  *[]domain.AlibabaAscpLogisticsConsignResendTopConsignPkgRequest `json:"consign_pkgs" required:"true" `
    /*
        feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号"12345678910-1234"     */
    Feature  *string `json:"feature,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsConsignResendRequest) SetTid(v string) *AlibabaAscpLogisticsConsignResendRequest {
    s.Tid = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignResendRequest) SetSubTids(v string) *AlibabaAscpLogisticsConsignResendRequest {
    s.SubTids = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignResendRequest) SetConsignPkgs(v []domain.AlibabaAscpLogisticsConsignResendTopConsignPkgRequest) *AlibabaAscpLogisticsConsignResendRequest {
    s.ConsignPkgs = &v
    return s
}
func (s *AlibabaAscpLogisticsConsignResendRequest) SetFeature(v string) *AlibabaAscpLogisticsConsignResendRequest {
    s.Feature = &v
    return s
}

func (req *AlibabaAscpLogisticsConsignResendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.SubTids != nil) {
        paramMap["sub_tids"] = *req.SubTids
    }
    if(req.ConsignPkgs != nil) {
        paramMap["consign_pkgs"] = util.ConvertStructList(*req.ConsignPkgs)
    }
    if(req.Feature != nil) {
        paramMap["feature"] = *req.Feature
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsConsignResendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}