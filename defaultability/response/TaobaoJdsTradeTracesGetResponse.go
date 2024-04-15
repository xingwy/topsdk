package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoJdsTradeTracesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        跟踪信息列表
    */
    Traces  []domain.TaobaoJdsTradeTracesGetTradeTrace `json:"traces,omitempty" `
    /*
        在订单全链路系统中的状态(比如是否存在)
    */
    UserStatus  string `json:"user_status,omitempty" `
}
