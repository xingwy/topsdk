package response

import (
)

type TaobaoTmcGroupDeleteResponse struct {

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
    IsSuccess  bool `json:"is_success,omitempty" `
}
