package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradePostageUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment
    */
    Trade  domain.TaobaoTradePostageUpdateTrade `json:"trade,omitempty" `
}
