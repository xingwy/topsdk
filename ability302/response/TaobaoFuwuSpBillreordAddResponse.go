package response

import (
)

type TaobaoFuwuSpBillreordAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回调用结果
    */
    AddResult  bool `json:"add_result,omitempty" `
}
