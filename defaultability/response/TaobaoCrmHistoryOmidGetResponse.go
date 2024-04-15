package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoCrmHistoryOmidGetResponse struct {

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
    Data  domain.TaobaoCrmHistoryOmidGetCrmPrivacyResponse `json:"data,omitempty" `
}
