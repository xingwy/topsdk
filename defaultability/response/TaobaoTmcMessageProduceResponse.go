package response

import (
)

type TaobaoTmcMessageProduceResponse struct {

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
    /*
        投递目标数
    */
    Total  int64 `json:"total,omitempty" `
    /*
        消息ID
    */
    MsgIds  []string `json:"msg_ids,omitempty" `
}
