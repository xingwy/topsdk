package response

import (
)

type AlibabaGpuAddSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回产品发布规则
    */
    AddProductRule  string `json:"add_product_rule,omitempty" `
}
