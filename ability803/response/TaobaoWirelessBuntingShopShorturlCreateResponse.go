package response

import (
)

type TaobaoWirelessBuntingShopShorturlCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        短链
    */
    Shorturl  string `json:"shorturl,omitempty" `
}
