package request

import (
        "topsdk/util"
    )

type AlibabaAscpLogisticsIdentcodeQueryRequest struct {
    /*
        根类目ID，"1101": "笔记本电脑",   "1512": "手机类目",   "50019780": "平板电脑"     */
    RootCatId  *string `json:"root_cat_id" required:"true" `
    /*
        识别码列表     */
    IdentCodeList  *[]string `json:"ident_code_list" required:"true" `
    /*
        识别码类型，SN，IMEI     */
    IdentType  *string `json:"ident_type" required:"true" `
    /*
        品牌ID，"11119": "Lenovo/联想",   "11656": "Asus/华硕",   "11813": "Huawei/华为",   "21660": "Hasee/神舟",   "21999": "MSI/微星",   "26683": "Dell/戴尔",   "26691": "Acer/宏碁",   "28247": "OPPO",   "30111": "Apple/苹果",   "31140": "HP/惠普",   "91621": "vivo",   "3506680": "MIUI/小米",   "184048021": "ThinkPad",   "590022244": "honor/荣耀",   "600184882": "OnePlus/一加",   "616784001": "MACHENIKE",   "639188075": "THUNDEROBOT",   "775486237": "MECHREVO/机械革命",   "1880225553": "ROG/玩家国度",   "7051840193": "iQOO(数码)"     */
    BrandId  *string `json:"brand_id" required:"true" `
    /*
        是否可用     */
    Available  *bool `json:"available,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsIdentcodeQueryRequest) SetRootCatId(v string) *AlibabaAscpLogisticsIdentcodeQueryRequest {
    s.RootCatId = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryRequest) SetIdentCodeList(v []string) *AlibabaAscpLogisticsIdentcodeQueryRequest {
    s.IdentCodeList = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryRequest) SetIdentType(v string) *AlibabaAscpLogisticsIdentcodeQueryRequest {
    s.IdentType = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryRequest) SetBrandId(v string) *AlibabaAscpLogisticsIdentcodeQueryRequest {
    s.BrandId = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeQueryRequest) SetAvailable(v bool) *AlibabaAscpLogisticsIdentcodeQueryRequest {
    s.Available = &v
    return s
}

func (req *AlibabaAscpLogisticsIdentcodeQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RootCatId != nil) {
        paramMap["root_cat_id"] = *req.RootCatId
    }
    if(req.IdentCodeList != nil) {
        paramMap["ident_code_list"] = util.ConvertBasicList(*req.IdentCodeList)
    }
    if(req.IdentType != nil) {
        paramMap["ident_type"] = *req.IdentType
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    if(req.Available != nil) {
        paramMap["available"] = *req.Available
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsIdentcodeQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}