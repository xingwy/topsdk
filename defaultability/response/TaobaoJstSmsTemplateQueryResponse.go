package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoJstSmsTemplateQueryResponse struct {

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
        返回结果
    */
    Module  domain.TaobaoJstSmsTemplateQueryAccessBaseDTO `json:"module,omitempty" `
    /*
        错误消息
    */
    Message  string `json:"message,omitempty" `
}
