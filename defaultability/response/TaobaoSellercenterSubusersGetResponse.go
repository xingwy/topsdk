package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoSellercenterSubusersGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
    */
    Subusers  []domain.TaobaoSellercenterSubusersGetSubUserInfo `json:"subusers,omitempty" `
}
