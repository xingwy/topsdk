package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemsOnsaleGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索到符合条件的结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段
    */
    Items  []domain.TaobaoItemsOnsaleGetItem `json:"items,omitempty" `
}
