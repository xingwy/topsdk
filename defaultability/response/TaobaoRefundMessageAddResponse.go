package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoRefundMessageAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        退款信息。包含id和created
    */
    RefundMessage  domain.TaobaoRefundMessageAddRefundMessage `json:"refund_message,omitempty" `
}
