package response

import (
)

type TmallItemAddSimpleschemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回发布商品的规则文档
    */
    Result  string `json:"result,omitempty" `
}
