package response

import (
    "topsdk/ability236/domain"
)

type CainiaoWaybillprintClientupdateGetconfigResponse struct {

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
    Result  domain.CainiaoWaybillprintClientupdateGetconfigUpdateConfigTopResult `json:"result,omitempty" `
}
