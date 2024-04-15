package response

import (
)

type TaobaoHttpdnsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        HTTP DNS配置信息
    */
    Result  string `json:"result,omitempty" `
}
