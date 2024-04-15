package response

import (
)

type TaobaoJstSmsMessageDirectBatchsendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        短信回执码
    */
    Module  string `json:"module,omitempty" `
}
