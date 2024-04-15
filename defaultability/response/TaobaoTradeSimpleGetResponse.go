package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoTradeSimpleGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        交易主订单信息
    */
    Trade  domain.TaobaoTradeSimpleGetTrade `json:"trade,omitempty" `
}
