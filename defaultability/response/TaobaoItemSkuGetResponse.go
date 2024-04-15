package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemSkuGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        Sku
    */
    Sku  domain.TaobaoItemSkuGetSku `json:"sku,omitempty" `
}
