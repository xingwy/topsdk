package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoRefundsApplyGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索到的交易信息总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        搜索到的退款信息列表
    */
    Refunds  []domain.TaobaoRefundsApplyGetRefund `json:"refunds,omitempty" `
}
