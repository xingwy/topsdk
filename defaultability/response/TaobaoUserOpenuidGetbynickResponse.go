package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoUserOpenuidGetbynickResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        买家uid对象
    */
    OpenUids  []domain.TaobaoUserOpenuidGetbynickOpenUidInfo `json:"open_uids,omitempty" `
}
