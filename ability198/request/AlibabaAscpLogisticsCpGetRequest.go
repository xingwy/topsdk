package request

import (
        "topsdk/ability198/domain"
        "topsdk/util"
    )

type AlibabaAscpLogisticsCpGetRequest struct {
    /*
        请求体     */
    LogisticsResourceRequest  *domain.AlibabaAscpLogisticsCpGetLogisticsResourceRequest `json:"logistics_resource_request,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsCpGetRequest) SetLogisticsResourceRequest(v domain.AlibabaAscpLogisticsCpGetLogisticsResourceRequest) *AlibabaAscpLogisticsCpGetRequest {
    s.LogisticsResourceRequest = &v
    return s
}

func (req *AlibabaAscpLogisticsCpGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LogisticsResourceRequest != nil) {
        paramMap["logistics_resource_request"] = util.ConvertStruct(*req.LogisticsResourceRequest)
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsCpGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}