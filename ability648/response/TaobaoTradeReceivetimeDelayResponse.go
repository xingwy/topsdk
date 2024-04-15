package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeReceivetimeDelayResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新后的交易数据，只包括tid和modified两个字段。
    */
    Trade  domain.TaobaoTradeReceivetimeDelayTrade `json:"trade,omitempty" `
}
