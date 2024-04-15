package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoKfcKeywordSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        KFC 关键词过滤匹配结果
    */
    KfcSearchResult  domain.TaobaoKfcKeywordSearchKfcSearchResult `json:"kfc_search_result,omitempty" `
}
