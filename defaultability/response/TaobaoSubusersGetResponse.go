package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSubusersGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        子账号基本信息
    */
    Subaccounts  []domain.TaobaoSubusersGetSubAccountInfo `json:"subaccounts,omitempty" `
}
