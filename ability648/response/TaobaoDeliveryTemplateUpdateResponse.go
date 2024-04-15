package response

import (
)

type TaobaoDeliveryTemplateUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        表示修改是否成功
    */
    Complete  bool `json:"complete,omitempty" `
}
