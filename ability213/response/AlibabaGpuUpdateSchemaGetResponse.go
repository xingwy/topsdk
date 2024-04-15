package response

import (
)

type AlibabaGpuUpdateSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        参数产品ID对应的产品更新规则
    */
    UpdateProductRule  string `json:"update_product_rule,omitempty" `
}
