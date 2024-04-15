package response

import (
)

type TmallItemHscodeDetailGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回的计量单位和销售单位
    */
    Results  []string `json:"results,omitempty" `
}
