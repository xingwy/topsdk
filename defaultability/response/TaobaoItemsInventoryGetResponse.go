package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemsInventoryGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段
    */
    Items  []domain.TaobaoItemsInventoryGetItem `json:"items,omitempty" `
    /*
        搜索到符合条件的结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
