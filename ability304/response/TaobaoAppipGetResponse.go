package response

import (
)

type TaobaoAppipGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        ISV发起请求服务器IP
    */
    Ip  string `json:"ip,omitempty" `
}
