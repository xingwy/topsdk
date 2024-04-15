package response

import (
)

type TaobaoCrmGroupUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        分组修改是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
