package response

import (
    "topsdk/defaultability/domain"
)

type AlibabaAliqinFcDigitalsmsCreatetemplateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回值
    */
    Result  domain.AlibabaAliqinFcDigitalsmsCreatetemplateResult `json:"result,omitempty" `
}
