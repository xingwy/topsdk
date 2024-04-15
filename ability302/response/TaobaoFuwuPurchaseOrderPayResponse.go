package response

import (
)

type TaobaoFuwuPurchaseOrderPayResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        该url用于订单付款
    */
    Url  string `json:"url,omitempty" `
}
