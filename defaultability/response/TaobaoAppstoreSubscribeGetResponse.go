package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoAppstoreSubscribeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        用户订购信息
    */
    UserSubscribe  domain.TaobaoAppstoreSubscribeGetUserSubscribe `json:"user_subscribe,omitempty" `
}
