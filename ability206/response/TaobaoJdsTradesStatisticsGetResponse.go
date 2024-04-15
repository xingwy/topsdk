package response

import (
    "topsdk/ability206/domain"
)

type TaobaoJdsTradesStatisticsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订单状态计数值
    */
    StatusInfos  []domain.TaobaoJdsTradesStatisticsGetTradeStat `json:"status_infos,omitempty" `
}
