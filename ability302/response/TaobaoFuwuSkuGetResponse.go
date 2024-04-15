package response

import (
    "topsdk/ability302/domain"
)

type TaobaoFuwuSkuGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        内购服务及SKU详情
    */
    Result  domain.TaobaoFuwuSkuGetArticleViewResult `json:"result,omitempty" `
}
