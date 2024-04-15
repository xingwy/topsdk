package response

import (
)

type AlibabaItemEditSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品已有规则信息，XML格式.
    */
    Result  string `json:"result,omitempty" `
}
