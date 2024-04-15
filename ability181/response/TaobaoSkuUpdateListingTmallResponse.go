package response

import (
    "topsdk/ability181/domain"
)

type TaobaoSkuUpdateListingTmallResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        商品模型信息
    */
    Model  domain.TaobaoSkuUpdateListingTmallItem `json:"model,omitempty" `
}
