package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoTmcUserGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        开通的用户数据
    */
    TmcUser  domain.TaobaoTmcUserGetTmcUser `json:"tmc_user,omitempty" `
}
