package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSkusCustomGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        Sku对象，具体字段以fields决定
    */
    Skus  []domain.TaobaoSkusCustomGetSku `json:"skus,omitempty" `
}
