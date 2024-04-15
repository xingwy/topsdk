package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        子账号被所拥有的权限
    */
    SubuserPermission  domain.TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission `json:"subuser_permission,omitempty" `
}
