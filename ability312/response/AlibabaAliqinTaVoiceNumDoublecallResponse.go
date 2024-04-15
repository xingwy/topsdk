package response

import (
    "topsdk/ability312/domain"
)

type AlibabaAliqinTaVoiceNumDoublecallResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回参数
    */
    Result  domain.AlibabaAliqinTaVoiceNumDoublecallBizResult `json:"result,omitempty" `
}
