package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeOrderskuUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        只返回oid和modified
    */
    Order  domain.TaobaoTradeOrderskuUpdateOrder `json:"order,omitempty" `
}
