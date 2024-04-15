package response

import (
)

type TaobaoCrmMemberinfoUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        会员信息修改是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
