package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemSellerGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品详细信息
    */
    Item  domain.TaobaoItemSellerGetItem `json:"item,omitempty" `
}
