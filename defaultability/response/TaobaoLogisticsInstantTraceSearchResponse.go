package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoLogisticsInstantTraceSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    Result  domain.TaobaoLogisticsInstantTraceSearchResult `json:"result,omitempty" `
}
