package response

import (
    "topsdk/ability198/domain"
)

type AlibabaAscpLogisticsSellerWritelogisticsnodeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    Result  domain.AlibabaAscpLogisticsSellerWritelogisticsnodeBatchWriteLogisticsNodeTopResponse `json:"result,omitempty" `
}
