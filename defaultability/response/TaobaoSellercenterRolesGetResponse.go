package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSellercenterRolesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
    */
    Roles  []domain.TaobaoSellercenterRolesGetRole `json:"roles,omitempty" `
}
