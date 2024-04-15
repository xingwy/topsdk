package response

import (
)

type TaobaoJstSmsTemplateReportResponse struct {

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
    RSuccess  string `json:"r_success,omitempty" `
    /*
        请求结果
    */
    Module  bool `json:"module,omitempty" `
    /*
        错误信息
    */
    Message  string `json:"message,omitempty" `
}
