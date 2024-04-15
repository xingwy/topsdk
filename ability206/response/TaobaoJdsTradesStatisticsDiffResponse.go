package response

import (
)

type TaobaoJdsTradesStatisticsDiffResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        pre_status比post_status多的tid列表
    */
    Tids  []int64 `json:"tids,omitempty" `
    /*
        总记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
