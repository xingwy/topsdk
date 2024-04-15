package response

import (
)

type TaobaoCrmGradeSetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        true：成功 false：失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
