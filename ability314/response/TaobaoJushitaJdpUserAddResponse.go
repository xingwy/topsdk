package response

import (
)

type TaobaoJushitaJdpUserAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否添加成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
