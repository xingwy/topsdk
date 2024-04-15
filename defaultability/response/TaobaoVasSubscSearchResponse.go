package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoVasSubscSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订购关系对象
    */
    ArticleSubs  []domain.TaobaoVasSubscSearchArticleSub `json:"article_subs,omitempty" `
    /*
        总记录数
    */
    TotalItem  int64 `json:"total_item,omitempty" `
}
