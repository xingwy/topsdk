package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoCrmHistoryOuidGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果
    */
    Data  domain.TaobaoCrmHistoryOuidGetCrmPrivacyResponse `json:"data,omitempty" `
}
