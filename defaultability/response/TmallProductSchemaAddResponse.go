package response

import (
)

type TmallProductSchemaAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        新发产品结果
    */
    AddProductResult  string `json:"add_product_result,omitempty" `
}
