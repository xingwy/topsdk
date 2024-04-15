package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoRefundGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        退款详情
    */
    Refund  domain.TaobaoRefundGetRefund `json:"refund,omitempty" `
}
