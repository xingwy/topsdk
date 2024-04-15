package response

import (
    "topsdk/ability680/domain"
)

type TaobaoOpencrmMemberSellerFetchstatusResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        total
    */
    Total  int64 `json:"total,omitempty" `
    /*
        result
    */
    SellerCleanStatusList  []domain.TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto `json:"seller_clean_status_list,omitempty" `
}
