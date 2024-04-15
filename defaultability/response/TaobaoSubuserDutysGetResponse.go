package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSubuserDutysGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        职务信息
    */
    Dutys  []domain.TaobaoSubuserDutysGetDuty `json:"dutys,omitempty" `
}
