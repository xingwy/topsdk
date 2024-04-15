package response

import (
)

type TaobaoQimenTradeUserDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        modal
    */
    Modal  bool `json:"modal,omitempty" `
}
