package response

import (
)

type TmallProductSchemaUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
    */
    UpdateProductResult  string `json:"update_product_result,omitempty" `
}
