package response

import (
)

type TaobaoFuwuSpConfirmApplyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回的是服务市场的确认单ID
    */
    ApplyResult  int64 `json:"apply_result,omitempty" `
}
