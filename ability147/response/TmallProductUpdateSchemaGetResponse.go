package response

import (
)

type TmallProductUpdateSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        参数产品ID对产品的更新规则
    */
    UpdateProductSchema  string `json:"update_product_schema,omitempty" `
}
