package response

import (
)

type TaobaoCrmGroupDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
