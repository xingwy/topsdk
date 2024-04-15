package response

import (
)

type AlibabaGpuSchemaUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新产品的结果
    */
    UpdateProductResult  string `json:"update_product_result,omitempty" `
}
