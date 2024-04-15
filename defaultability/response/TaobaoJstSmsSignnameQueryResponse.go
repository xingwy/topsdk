package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoJstSmsSignnameQueryResponse struct {

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
        返回结果（signStatus： 0--待审核  1--通过 2--拒绝）
    */
    Module  domain.TaobaoJstSmsSignnameQueryQuerySmsSignDTO `json:"module,omitempty" `
    /*
        失败原因
    */
    Message  string `json:"message,omitempty" `
}
