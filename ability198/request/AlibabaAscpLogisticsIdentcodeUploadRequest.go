package request

import (
        "topsdk/util"
    )

type AlibabaAscpLogisticsIdentcodeUploadRequest struct {
    /*
        订单id     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        sn码列表，多个用逗号分隔     */
    CodeList  *[]string `json:"code_list" required:"true" `
}

func (s *AlibabaAscpLogisticsIdentcodeUploadRequest) SetOrderId(v int64) *AlibabaAscpLogisticsIdentcodeUploadRequest {
    s.OrderId = &v
    return s
}
func (s *AlibabaAscpLogisticsIdentcodeUploadRequest) SetCodeList(v []string) *AlibabaAscpLogisticsIdentcodeUploadRequest {
    s.CodeList = &v
    return s
}

func (req *AlibabaAscpLogisticsIdentcodeUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.CodeList != nil) {
        paramMap["code_list"] = util.ConvertBasicList(*req.CodeList)
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsIdentcodeUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}