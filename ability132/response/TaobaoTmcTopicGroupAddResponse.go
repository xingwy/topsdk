package response

import (
)

type TaobaoTmcTopicGroupAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        true
    */
    Result  bool `json:"result,omitempty" `
}
