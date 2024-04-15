package response

import (
)

type TmallItemCombineGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        results
    */
    Results  []string `json:"results,omitempty" `
}
