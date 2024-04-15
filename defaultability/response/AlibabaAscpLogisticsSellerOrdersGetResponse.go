package response

import (
    "topsdk/defaultability/domain"
)

type AlibabaAscpLogisticsSellerOrdersGetResponse struct {

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
    Result  domain.AlibabaAscpLogisticsSellerOrdersGetResultDTO `json:"result,omitempty" `
}
