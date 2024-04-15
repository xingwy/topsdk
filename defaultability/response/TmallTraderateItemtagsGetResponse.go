package response

import (
    "topsdk/defaultability/domain"
)

type TmallTraderateItemtagsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        标签列表
    */
    Tags  []domain.TmallTraderateItemtagsGetTmallRateTagDetail `json:"tags,omitempty" `
}
