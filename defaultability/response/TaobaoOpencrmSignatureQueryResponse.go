package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoOpencrmSignatureQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        短信签名列表
    */
    ResultList  []domain.TaobaoOpencrmSignatureQueryCrmShortMessageSignatureDo `json:"result_list,omitempty" `
}
