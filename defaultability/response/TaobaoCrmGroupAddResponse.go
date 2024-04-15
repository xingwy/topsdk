package response

import (
)

type TaobaoCrmGroupAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        新增分组的id
    */
    GroupId  int64 `json:"group_id,omitempty" `
    /*
        添加分组是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
