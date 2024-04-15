package response

import (
)

type TaobaoLogisticsConsignOrderCreateandsendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果描述
    */
    ResultDesc  string `json:"result_desc,omitempty" `
    /*
        是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        订单ID
    */
    OrderId  int64 `json:"order_id,omitempty" `
}
