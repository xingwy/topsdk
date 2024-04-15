package response

import (
    "topsdk/ability198/domain"
)

type AlibabaAscpLogisticsCpGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回值
    */
    Result  domain.AlibabaAscpLogisticsCpGetResultDTO `json:"result,omitempty" `
}
