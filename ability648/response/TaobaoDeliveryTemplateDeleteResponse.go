package response

import (
)

type TaobaoDeliveryTemplateDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        表示删除成功还是失败
    */
    Complete  bool `json:"complete,omitempty" `
}
