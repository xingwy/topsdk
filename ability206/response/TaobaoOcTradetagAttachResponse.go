package response

import (
)

type TaobaoOcTradetagAttachResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作成功或者操作失败
    */
    Result  bool `json:"result,omitempty" `
}
