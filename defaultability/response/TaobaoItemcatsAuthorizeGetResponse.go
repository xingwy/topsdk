package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemcatsAuthorizeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        里面有3个数组：Brand[]品牌列表,ItemCat[] 类目列表XinpinItemCat[] 针对于C卖家新品类目列表
    */
    SellerAuthorize  domain.TaobaoItemcatsAuthorizeGetSellerAuthorize `json:"seller_authorize,omitempty" `
}
