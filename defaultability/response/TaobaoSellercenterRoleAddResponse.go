package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSellercenterRoleAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        子账号角色
    */
    Role  domain.TaobaoSellercenterRoleAddRole `json:"role,omitempty" `
}
