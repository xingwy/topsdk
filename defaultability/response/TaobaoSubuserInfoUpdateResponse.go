package response

import (
)

type TaobaoSubuserInfoUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作是否成功 true:操作成功; false:操作失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
