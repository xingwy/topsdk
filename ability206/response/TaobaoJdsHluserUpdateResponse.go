package response

import (
)

type TaobaoJdsHluserUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功
    */
    Result  bool `json:"result,omitempty" `
}
