package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeShippingaddressUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        交易结构
    */
    Trade  domain.TaobaoTradeShippingaddressUpdateTrade `json:"trade,omitempty" `
}
