package response

import (
)

type TaobaoJstSmsTemplateModifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误码
    */
    RCode  string `json:"r_code,omitempty" `
    /*
        请求成功
    */
    RSuccess  bool `json:"r_success,omitempty" `
    /*
        修改成功
    */
    Module  bool `json:"module,omitempty" `
    /*
        错误信息
    */
    Message  string `json:"message,omitempty" `
}
