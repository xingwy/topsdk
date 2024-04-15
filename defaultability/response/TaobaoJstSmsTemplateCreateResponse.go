package response

import (
)

type TaobaoJstSmsTemplateCreateResponse struct {

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
        模板CODE
    */
    Module  string `json:"module,omitempty" `
    /*
        失败原因
    */
    Message  string `json:"message,omitempty" `
}
