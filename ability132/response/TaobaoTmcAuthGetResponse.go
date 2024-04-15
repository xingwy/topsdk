package response

import (
)

type TaobaoTmcAuthGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  string `json:"result,omitempty" `
}
