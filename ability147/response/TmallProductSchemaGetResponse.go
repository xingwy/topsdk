package response

import (
)

type TmallProductSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品信息数据。schema形式
    */
    GetProductResult  string `json:"get_product_result,omitempty" `
}
