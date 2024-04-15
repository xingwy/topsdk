package response

import (
    "topsdk/ability200/domain"
)

type TaobaoRefundDetailGetResponse struct {

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
    Detail  domain.TaobaoRefundDetailGetRefundDetail `json:"detail,omitempty" `
}
