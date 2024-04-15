package response

import (
    "topsdk/ability145/domain"
)

type TaobaoTradeSellermemosGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        备注列表
    */
    Memos  []domain.TaobaoTradeSellermemosGetOrderSellerMemo `json:"memos,omitempty" `
}
