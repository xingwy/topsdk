package response

import (
    "topsdk/ability315/domain"
)

type TaobaoRdsDbGetdbResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  domain.TaobaoRdsDbGetdbResultSet `json:"result,omitempty" `
}
