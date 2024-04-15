package response

import (
)

type TaobaoOcTradesBytagGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        打了该标签的订单编号列表
    */
    Tids  []int64 `json:"tids,omitempty" `
    /*
        总数
    */
    Totals  int64 `json:"totals,omitempty" `
}
