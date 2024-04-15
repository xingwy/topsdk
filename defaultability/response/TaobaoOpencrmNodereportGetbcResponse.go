package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoOpencrmNodereportGetbcResponse struct {

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
    Result  domain.TaobaoOpencrmNodereportGetbcNodeExecuteReportDto `json:"result,omitempty" `
}
