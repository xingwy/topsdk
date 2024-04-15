package response

import (
)

type TmallProductAddSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回发布产品的规则文档
    */
    AddProductRule  string `json:"add_product_rule,omitempty" `
}
