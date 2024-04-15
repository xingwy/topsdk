package response

import (
)

type TaobaoJstSmsSignnameCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误CODE
    */
    RCode  string `json:"r_code,omitempty" `
    /*
        请求是否成功
    */
    RSuccess  bool `json:"r_success,omitempty" `
    /*
        请求成功
    */
    Module  bool `json:"module,omitempty" `
    /*
        失败原因
    */
    Message  string `json:"message,omitempty" `
}
