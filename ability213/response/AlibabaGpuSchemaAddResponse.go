package response

import (
)

type AlibabaGpuSchemaAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品发布的结果
    */
    AddProductResult  string `json:"add_product_result,omitempty" `
}
