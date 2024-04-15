package response

import (
    "topsdk/ability145/domain"
)

type TaobaoTradeSellerflagsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        插旗列表
    */
    Flags  []domain.TaobaoTradeSellerflagsGetSellerFlag `json:"flags,omitempty" `
}
