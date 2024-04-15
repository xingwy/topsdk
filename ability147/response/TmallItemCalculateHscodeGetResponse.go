package response

import (
)

type TmallItemCalculateHscodeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        算法返回预测的hscode数据
    */
    Results  []string `json:"results,omitempty" `
}
