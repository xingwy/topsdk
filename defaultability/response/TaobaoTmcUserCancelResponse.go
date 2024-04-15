package response

import (
)

type TaobaoTmcUserCancelResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功,如果为false并且没有错误码，表示删除的用户不存在。
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
