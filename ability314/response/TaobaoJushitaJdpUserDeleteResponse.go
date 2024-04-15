package response

import (
)

type TaobaoJushitaJdpUserDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否删除成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
