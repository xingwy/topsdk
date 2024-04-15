package response

import (
    "topsdk/ability198/domain"
)

type AlibabaAscpLogisticsIdentcodeUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        请求结果
    */
    Result  domain.AlibabaAscpLogisticsIdentcodeUploadMptResult `json:"result,omitempty" `
}
