package response

import (
)

type TaobaoCrmMembersGroupsBatchdeletePrivyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        删除是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
