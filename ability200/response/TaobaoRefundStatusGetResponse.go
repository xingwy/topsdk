package response

import (
    "topsdk/ability200/domain"
)

type TaobaoRefundStatusGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        出参对象
    */
    ResultPackage  domain.TaobaoRefundStatusGetResultSet `json:"result_package,omitempty" `
}
