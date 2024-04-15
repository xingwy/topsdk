package response

import (
)

type TaobaoJushitaJmsTopicsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误信息
    */
    ResultMessage  string `json:"result_message,omitempty" `
    /*
        topic列表
    */
    Results  []string `json:"results,omitempty" `
    /*
        错误码
    */
    ResultCode  string `json:"result_code,omitempty" `
}
